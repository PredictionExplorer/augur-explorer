/*
 * setuid-root launcher for jailed Codex MCP (cosmic1).
 *
 * Why this exists:
 * - cgprod must run nsjail with euid=0 for namespace mounts (--disable_clone_newuser).
 * - setuid on a shell script is ignored by Linux.
 * - execve() of a non-setuid nsjail drops elevated euid, so a wrapper must exec a
 *   setuid nsjail binary (typically /usr/bin/nsjail, not under a nosuid mount).
 *
 * Build + install (as root) — install under /usr/bin, NOT /usr/local (often nosuid):
 *   gcc -O2 -Wall -Wextra -o /usr/bin/codex-nsjail-launcher codex-nsjail-launcher.c
 *   chown root:cgprod /usr/bin/codex-nsjail-launcher
 *   chmod 4750 /usr/bin/codex-nsjail-launcher
 *
 *   cp /usr/local/bin/nsjail /usr/bin/nsjail
 *   chown root:cgprod /usr/bin/nsjail
 *   chmod 4750 /usr/bin/nsjail
 *
 * If strace still shows geteuid()=1008 before clone3 EPERM, use file caps instead:
 *   setcap cap_sys_admin+ep /usr/bin/nsjail
 *   chmod 750 /usr/bin/nsjail
 *
 *   mkdir -p /var/lib/nsjail-aijail && chown aijail:aijail /var/lib/nsjail-aijail && chmod 700 ...
 *
 *   CODEX_MCP_COMMAND=/usr/bin/codex-nsjail-launcher
 *   CODEX_MCP_ARGS=mcp-server
 *
 * Verify as cgprod:
 *   strace -e getuid,geteuid,clone3 /usr/bin/codex-nsjail-launcher mcp-server 2>&1 | head -30
 *
 * Fallback: sudo (see faq-bot.env.example Option B).
 */
#define _GNU_SOURCE
#include <errno.h>
#include <grp.h>
#include <pwd.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#ifndef ALLOWED_USER
#define ALLOWED_USER "cgprod"
#endif

#ifndef NSJAIL_BIN
#define NSJAIL_BIN "/usr/bin/nsjail"
#endif

#ifndef CODEX_BIN
#define CODEX_BIN "/home/aijail/.local/bin/codex"
#endif

#ifndef NODE_DIR
#define NODE_DIR "/usr/local/node-v18.20.5-linux-x64"
#endif

#ifndef NSJAIL_WORK
#define NSJAIL_WORK "/var/lib/nsjail-aijail"
#endif

static uid_t allowed_uid(void) {
    struct passwd *pw = getpwnam(ALLOWED_USER);
    if (!pw) {
        fprintf(stderr, "codex-nsjail-launcher: unknown user %s\n", ALLOWED_USER);
        exit(1);
    }
    return pw->pw_uid;
}

static gid_t aijail_gid(void) {
    struct passwd *pw = getpwnam("aijail");
    if (!pw) {
        fprintf(stderr, "codex-nsjail-launcher: unknown user aijail\n");
        exit(1);
    }
    return pw->pw_gid;
}

static uid_t aijail_uid(void) {
    struct passwd *pw = getpwnam("aijail");
    if (!pw) {
        fprintf(stderr, "codex-nsjail-launcher: unknown user aijail\n");
        exit(1);
    }
    return pw->pw_uid;
}

static bool argv_is_mcp_server(int argc, char *argv[]) {
    return argc == 2 && strcmp(argv[1], "mcp-server") == 0;
}

int main(int argc, char *argv[]) {
    if (getuid() != allowed_uid()) {
        fprintf(stderr, "codex-nsjail-launcher: permission denied\n");
        exit(1);
    }
    if (!argv_is_mcp_server(argc, argv)) {
        fprintf(stderr, "codex-nsjail-launcher: only 'mcp-server' is allowed\n");
        exit(1);
    }
    if (access(NSJAIL_BIN, X_OK) != 0) {
        fprintf(stderr, "codex-nsjail-launcher: nsjail not executable: %s (%s)\n",
                NSJAIL_BIN, strerror(errno));
        exit(1);
    }

    char uid_buf[32];
    char gid_buf[32];
    char path_val[512];
    char path_env[520];
    snprintf(uid_buf, sizeof uid_buf, "%u", (unsigned)aijail_uid());
    snprintf(gid_buf, sizeof gid_buf, "%u", (unsigned)aijail_gid());
    snprintf(path_val, sizeof path_val,
             "/home/aijail/.local/bin:%s/bin:/usr/local/bin:/usr/bin:/bin", NODE_DIR);
    snprintf(path_env, sizeof path_env, "PATH=%s", path_val);

    setenv("XDG_RUNTIME_DIR", NSJAIL_WORK, 1);
    setenv("HOME", "/home/aijail", 1);
    setenv("CODEX_HOME", "/home/aijail/.codex", 1);
    setenv("PATH", path_val, 1);

    char *nsjail_argv[] = {
        (char *)"nsjail",
        (char *)"--mode", (char *)"o",
        (char *)"--hostname", (char *)"codexjail",
        (char *)"--disable_clone_newuser",
        (char *)"--disable_clone_newpid",
        (char *)"--disable_clone_newnet",
        (char *)"--time_limit", (char *)"0",
        (char *)"--rlimit_as", (char *)"2048",
        (char *)"--rlimit_cpu", (char *)"300",
        (char *)"--rlimit_fsize", (char *)"32",
        (char *)"--rlimit_nofile", (char *)"512",
        (char *)"--rlimit_nproc", (char *)"512",
        (char *)"--user", uid_buf,
        (char *)"--group", gid_buf,
        (char *)"--cwd", (char *)"/home/aijail",
        (char *)"--env", (char *)"HOME=/home/aijail",
        (char *)"--env", (char *)"CODEX_HOME=/home/aijail/.codex",
        (char *)"--env", path_env,
        (char *)"--mount", (char *)"/dev:/dev",
        (char *)"--tmpfsmount", (char *)"/tmp:size=256,mode=1777",
        (char *)"--tmpfsmount", (char *)"/dev/shm:size=64,mode=1777",
        (char *)"--bindmount_ro", (char *)"/usr",
        (char *)"--bindmount_ro", (char *)"/bin",
        (char *)"--bindmount_ro", (char *)"/lib",
        (char *)"--bindmount_ro", (char *)"/lib64",
        (char *)"--bindmount_ro", (char *)"/etc/ssl",
        (char *)"--bindmount_ro", (char *)"/etc/resolv.conf",
        (char *)"--bindmount_ro", (char *)"/etc/hosts",
        (char *)"--bindmount_ro", (char *)"/etc/nsswitch.conf",
        (char *)"--bindmount", (char *)NODE_DIR,
        (char *)"--bindmount", (char *)"/home/aijail",
        (char *)"--",
        (char *)CODEX_BIN,
        (char *)"mcp-server",
        NULL,
    };

    execv(NSJAIL_BIN, nsjail_argv);
    fprintf(stderr, "codex-nsjail-launcher: exec %s failed: %s\n", NSJAIL_BIN, strerror(errno));
    return 1;
}
