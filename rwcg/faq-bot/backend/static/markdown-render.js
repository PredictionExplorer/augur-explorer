/**
 * Lightweight markdown renderer for the FAQ test UI.
 * Mirrors app/components/MarkdownRenderer.tsx — no external dependencies.
 * Supports: **bold**, `code`, *italic*, bullet lists, numbered lists, blank lines.
 */
(function (global) {
  'use strict';

  function renderInline(text) {
    const parts = text.split(/(\*\*[^*]+\*\*|`[^`]+`|\*[^*]+\*)/);
    return parts.map(function (part) {
      if (/^\*\*[^*]+\*\*$/.test(part)) {
        const strong = document.createElement('strong');
        strong.textContent = part.slice(2, -2);
        return strong;
      }
      if (/^`[^`]+`$/.test(part)) {
        const code = document.createElement('code');
        code.textContent = part.slice(1, -1);
        return code;
      }
      if (/^\*[^*]+\*$/.test(part)) {
        const em = document.createElement('em');
        em.textContent = part.slice(1, -1);
        return em;
      }
      return document.createTextNode(part);
    });
  }

  function appendInline(parent, text) {
    renderInline(text).forEach(function (node) {
      parent.appendChild(node);
    });
  }

  function renderMarkdownTo(container, content) {
    container.replaceChildren();
    const root = document.createElement('div');
    root.className = 'md-root';

    const lines = content.split('\n');
    let i = 0;

    while (i < lines.length) {
      const line = lines[i];

      if (line.trim() === '') {
        const spacer = document.createElement('div');
        spacer.className = 'md-spacer';
        root.appendChild(spacer);
        i++;
        continue;
      }

      if (/^\*\*[^*]+\*\*$/.test(line.trim())) {
        const p = document.createElement('p');
        p.className = 'md-heading';
        p.textContent = line.trim().replace(/^\*\*|\*\*$/g, '');
        root.appendChild(p);
        i++;
        continue;
      }

      if (/^(\s*[-*])\s+/.test(line)) {
        const ul = document.createElement('ul');
        ul.className = 'md-list';
        while (i < lines.length && /^(\s*[-*])\s+/.test(lines[i])) {
          const li = document.createElement('li');
          appendInline(li, lines[i].replace(/^\s*[-*]\s+/, ''));
          ul.appendChild(li);
          i++;
        }
        root.appendChild(ul);
        continue;
      }

      if (/^\d+\.\s+/.test(line)) {
        const ol = document.createElement('ol');
        ol.className = 'md-list';
        while (i < lines.length && /^\d+\.\s+/.test(lines[i])) {
          const li = document.createElement('li');
          appendInline(li, lines[i].replace(/^\d+\.\s+/, ''));
          ol.appendChild(li);
          i++;
        }
        root.appendChild(ol);
        continue;
      }

      const p = document.createElement('p');
      p.className = 'md-paragraph';
      appendInline(p, line);
      root.appendChild(p);
      i++;
    }

    container.appendChild(root);
  }

  global.renderMarkdownTo = renderMarkdownTo;
})(window);
