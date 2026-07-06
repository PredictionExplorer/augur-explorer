'use client'

import React, { ReactNode } from 'react'

interface MarkdownRendererProps {
  content: string
}

/**
 * Lightweight markdown renderer — no external dependencies.
 * Supports: **bold**, `code`, *italic*, bullet lists, numbered lists, headings, blank lines.
 */
export default function MarkdownRenderer({ content }: MarkdownRendererProps) {
  const lines = content.split('\n')
  const elements: ReactNode[] = []
  let i = 0

  while (i < lines.length) {
    const line = lines[i]

    // Blank line → spacer
    if (line.trim() === '') {
      elements.push(<div key={i} style={{ height: '0.4rem' }} />)
      i++
      continue
    }

    // Standalone **Heading** line
    if (/^\*\*[^*]+\*\*$/.test(line.trim())) {
      const text = line.trim().replace(/^\*\*|\*\*$/g, '')
      elements.push(
        <p key={i} style={{ fontWeight: 700, marginTop: '0.75rem', marginBottom: '0.1rem' }}>
          {text}
        </p>
      )
      i++
      continue
    }

    // Bullet list
    if (/^(\s*[-*])\s+/.test(line)) {
      const listItems: ReactNode[] = []
      while (i < lines.length && /^(\s*[-*])\s+/.test(lines[i])) {
        const itemText = lines[i].replace(/^\s*[-*]\s+/, '')
        listItems.push(<li key={i}>{renderInline(itemText)}</li>)
        i++
      }
      elements.push(
        <ul key={`ul-${i}`} style={{ paddingLeft: '1.25rem', margin: '0.2rem 0' }}>
          {listItems}
        </ul>
      )
      continue
    }

    // Numbered list
    if (/^\d+\.\s+/.test(line)) {
      const listItems: ReactNode[] = []
      while (i < lines.length && /^\d+\.\s+/.test(lines[i])) {
        const itemText = lines[i].replace(/^\d+\.\s+/, '')
        listItems.push(<li key={i}>{renderInline(itemText)}</li>)
        i++
      }
      elements.push(
        <ol key={`ol-${i}`} style={{ paddingLeft: '1.25rem', margin: '0.2rem 0' }}>
          {listItems}
        </ol>
      )
      continue
    }

    // Normal paragraph
    elements.push(
      <p key={i} style={{ margin: '0.2rem 0', lineHeight: 1.6 }}>
        {renderInline(line)}
      </p>
    )
    i++
  }

  return <div style={{ fontSize: '0.95rem' }}>{elements}</div>
}

/**
 * Render inline markdown: **bold**, `code`, *italic*, plain text
 */
function renderInline(text: string): ReactNode[] {
  const parts = text.split(/(\*\*[^*]+\*\*|`[^`]+`|\*[^*]+\*)/)
  return parts.map((part, idx) => {
    if (/^\*\*[^*]+\*\*$/.test(part)) {
      return <strong key={idx}>{part.slice(2, -2)}</strong>
    }
    if (/^`[^`]+`$/.test(part)) {
      return (
        <code
          key={idx}
          style={{
            background: 'rgba(0,0,0,0.08)',
            borderRadius: '3px',
            padding: '1px 5px',
            fontFamily: 'monospace',
            fontSize: '0.88em',
          }}
        >
          {part.slice(1, -1)}
        </code>
      )
    }
    if (/^\*[^*]+\*$/.test(part)) {
      return <em key={idx}>{part.slice(1, -1)}</em>
    }
    return <span key={idx}>{part}</span>
  })
}
