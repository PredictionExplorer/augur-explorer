"""
Random Walk NFT raster + video generation (ported from randomwalknft-back-end-2).

Requires: Pillow, numpy, opencv-python-headless, ffmpeg on PATH.

Set RWALK_IMG_OUTPUT_DIR to the directory that maps to GET /images/randomwalk/ on websrv
(i.e. the `randomwalk` folder inside NFT_ASSETS_ROOT when the static root is the parent of `randomwalk/`).
"""

from PIL import Image, ImageDraw
import cv2
import hashlib
import logging
import numpy as np
import os
import subprocess

VIDEO_FPS = 60


def _output_dir() -> str:
    d = os.environ.get("RWALK_IMG_OUTPUT_DIR", "").strip()
    if not d:
        raise RuntimeError("RWALK_IMG_OUTPUT_DIR is not set (absolute path to .../randomwalk asset folder)")
    return os.path.abspath(d)

def random_generator(init_seed):
    '''Generate random bits.'''
    if init_seed is None:
        raise ValueError("init_seed cannot be None")
    if init_seed.startswith('0x'):
        init_seed = init_seed[2:]
    init_seed = bytes.fromhex(init_seed)

    seed = init_seed

    while True:
        m = hashlib.sha3_256()
        m.update(init_seed)
        m.update(seed)
        seed = m.digest()
        for b in seed:
            for i in range(8):
                yield (b >> i) & 1


def create_media_helper(file_name, seed, background_color):
    '''Generate a PNG image and 2 MP4 videos.'''

    OUTPUT_DIR = _output_dir()
    os.makedirs(OUTPUT_DIR, exist_ok=True)

    file_name = file_name + '_' + background_color

    gen = random_generator(seed)

    horizontal_steps = []
    vertical_steps = []

    vert = 1500
    target_size = (int(vert * 1.6), vert)

    x, y = 0, 0
    min_x, max_x, min_y, max_y = 0, 0, 0, 0
    while True:
        a, b = next(gen), next(gen)
        if (a, b) == (0, 0):
            x += 1
            horizontal_steps.append(1)
            vertical_steps.append(0)
        elif (a, b) == (0, 1):
            x -= 1
            horizontal_steps.append(-1)
            vertical_steps.append(0)
        elif (a, b) == (1, 0):
            y += 1
            horizontal_steps.append(0)
            vertical_steps.append(1)
        elif (a, b) == (1, 1):
            y -= 1
            horizontal_steps.append(0)
            vertical_steps.append(-1)

        min_x = min(min_x, x)
        max_x = max(max_x, x)

        min_y = min(min_y, y)
        max_y = max(max_y, y)

        x_range = max_x - min_x
        y_range = max_y - min_y

        longer_range = max(x_range, y_range)
        shorter_range = min(x_range, y_range)

        if longer_range >= target_size[0] or shorter_range >= target_size[1]:
            break

    if x_range < y_range:
        vertical_steps, horizontal_steps = horizontal_steps, vertical_steps
        min_x, max_x, min_y, max_y = min_y, max_y, min_x, max_x
    x_range = max_x - min_x
    y_range = max_y - min_y

    num_steps = len(horizontal_steps)
    logging.info(f"Number of steps in walk: {num_steps}")

    origin = np.zeros((1, 2))

    steps = np.stack((horizontal_steps, vertical_steps), axis=1)

    path = np.concatenate([origin, steps]).cumsum(0)

    def random_color_1ch(num_steps):
        '''Geenrate colors for 1 channel.'''
        cur = 0
        result = []
        for _ in range(num_steps):
            cur += 1 if next(gen) == 1 else -1
            result.append(cur)
        lowest = min(result)
        highest = max(result)
        for i in range(len(result)):
            result[i] = (result[i] - lowest) / (highest - lowest)
        return result

    c1 = random_color_1ch(num_steps + 1)
    c2 = random_color_1ch(num_steps + 1)
    c3 = random_color_1ch(num_steps + 1)

    C = np.array(list(zip(c1, c2, c3)))

    x_center = (min_x + max_x) / 2
    y_center = (min_y + max_y) / 2

    BORDER_PERCENT = 0.03
    border = int(target_size[1] * BORDER_PERCENT)

    final_size = tuple(x + 2 * border for x in target_size)
    im = Image.new('RGB', final_size, background_color)
    draw = ImageDraw.Draw(im)

    for i, step in enumerate(path):
        x, y = step
        x = int(x - x_center + target_size[0] / 2) + border
        y = int(y - y_center + target_size[1] / 2) + border
        draw.point((x, y), fill=tuple(int(x * 255) for x in C[i]))

    png_path = os.path.join(OUTPUT_DIR, f"{file_name}.png")
    jpg_path = os.path.join(OUTPUT_DIR, f"{file_name}_thumb.jpg")

    im.save(png_path, "PNG")
    logging.info(f"{png_path} saved.")

    width, height = im.size
    new_size = (width // 4, height // 4)
    im = im.resize(new_size, resample=Image.LANCZOS)

    im.save(jpg_path, quality=95)
    logging.info(f"{jpg_path} saved.")

    def generate_video(walkers, label):
        '''Generate a video. num_walkers is the number of starting points for
        the random walk.'''
        images = []
        im = Image.new('RGB', final_size, background_color)
        images.append(im)
        draw = ImageDraw.Draw(im)

        visited_coordinates = np.zeros(final_size)

        jump = len(path) // 600


        visited_index = set()

        def advance_walker(walker_num, draw):
            index, direction, active = walkers[walker_num]
            if not active:
                return 0
            new_index = index + direction

            if new_index in visited_index or new_index < 0 or new_index >= len(path):
                walkers[walker_num] = (index, direction, False)
                return 0

            walkers[walker_num] = (new_index, direction, True)

            visited_index.add(new_index)

            x, y = path[new_index]
            x = int(x - x_center + target_size[0] / 2) + border
            y = int(y - y_center + target_size[1] / 2) + border

            if visited_coordinates[x][y] < new_index:
                draw.point((x, y), fill=tuple(int(x * 255) for x in C[new_index]))
                visited_coordinates[x][y] = new_index
            return 1

        walker_num = 0
        since_frame = 0
        while any(x[-1] for x in walkers):
            since_frame += advance_walker(walker_num % len(walkers), draw)
            walker_num += 1
            if since_frame >= jump:
                images.append(im)
                im = im.copy()
                draw = ImageDraw.Draw(im)
                since_frame = 0
            else:
                since_frame += 1

        def add_holds(images):
            '''Add a few seconds before and at the end of the video.'''
            result = []
            INIT_HOLD_SECONDS = 0.3
            blank = Image.new('RGB', final_size, background_color)

            for _ in range(int(INIT_HOLD_SECONDS * VIDEO_FPS)):
                result.append(blank)

            result.extend(images)

            END_HOLD_SECONDS = 2
            for _ in range(END_HOLD_SECONDS * VIDEO_FPS):
                result.append(images[-1])
            return result

        images = add_holds(images)

        out_path = os.path.join(OUTPUT_DIR, f'{file_name}_{label}.mp4')
        raw_path = os.path.join(OUTPUT_DIR, f'{file_name}_{label}_raw.mp4')
        out = cv2.VideoWriter(raw_path, cv2.VideoWriter_fourcc(*'MP4V'), VIDEO_FPS, final_size)

        for i in range(len(images)):
            cv_img = cv2.cvtColor(np.array(images[i]), cv2.COLOR_RGB2BGR)
            out.write(cv_img)

        out.release()
        logging.info("%s intermediate saved.", raw_path)

        subprocess.run(
            ["ffmpeg", "-y", "-i", raw_path, "-vcodec", "libx264", out_path],
            check=True,
        )
        os.remove(raw_path)
        logging.info("%s saved.", out_path)

    # Generate video with 1 starting point
    walkers = [(-1, 1, True)]
    generate_video(walkers, "single")

    # Generate video with 3 starting points
    num_walkers = 3
    walkers = []
    for i in range(num_walkers):
        k = i / num_walkers + (1 / (num_walkers * 2))
        c = int(k * len(path))
        walkers.append((c - 1, 1, True))
        walkers.append((c, -1, True))
    generate_video(walkers, "triple")

def create_media(file_name, seed):
    create_media_helper(file_name, seed, 'white')
    create_media_helper(file_name, seed, 'black')
