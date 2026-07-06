'use client'

import { useState, useRef, useEffect } from 'react'
import styles from './ChatInterface.module.css'
import MarkdownRenderer from './MarkdownRenderer'

interface Message {
  id: string
  text: string
  sender: 'user' | 'bot'
  timestamp: Date
  sources?: string[]
}

export default function ChatInterface() {
  const [messages, setMessages] = useState<Message[]>([
    {
      id: '1',
      text: "Hello! I'm the Cosmic Signature FAQ bot. I answer questions using our curated knowledge base (Haystack) and Codex AI. Ask me anything — from how to bid to smart contract details!",
      sender: 'bot',
      timestamp: new Date(),
    },
  ])
  const [inputValue, setInputValue] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const [sessionId, setSessionId] = useState<string | null>(null)
  const messagesEndRef = useRef<HTMLDivElement>(null)

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }

  useEffect(() => {
    scrollToBottom()
  }, [messages])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!inputValue.trim() || isLoading) return

    const userMessage: Message = {
      id: Date.now().toString(),
      text: inputValue,
      sender: 'user',
      timestamp: new Date(),
    }

    const question = inputValue
    setMessages((prev) => [...prev, userMessage])
    setInputValue('')
    setIsLoading(true)

    try {
      const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL || 'http://localhost:8000'
      const response = await fetch(`${backendUrl}/api/query`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ question, session_id: sessionId }),
      })

      const data = await response.json()

      if (!response.ok) {
        const detail = data.detail
        const component =
          typeof detail === 'object' && detail?.component ? detail.component : 'unknown'
        const message =
          typeof detail === 'object' && detail?.error
            ? detail.error
            : typeof detail === 'string'
              ? detail
              : 'Failed to get response from bot'
        throw new Error(`[${component}] ${message}`)
      }

      if (data.session_id) {
        setSessionId(data.session_id)
      }

      const botMessage: Message = {
        id: (Date.now() + 1).toString(),
        text: data.answer || 'I apologize, but I could not find an answer to your question.',
        sender: 'bot',
        timestamp: new Date(),
        sources: data.sources,
      }

      setMessages((prev) => [...prev, botMessage])
    } catch (error) {
      console.error('Error:', error)
      const errText =
        error instanceof Error
          ? error.message
          : 'Sorry, I encountered an error. Please make sure the backend server is running and try again.'
      const errorMessage: Message = {
        id: (Date.now() + 1).toString(),
        text: `⚠️ ${errText}`,
        sender: 'bot',
        timestamp: new Date(),
      }
      setMessages((prev) => [...prev, errorMessage])
    } finally {
      setIsLoading(false)
    }
  }

  const suggestedQuestions = [
    'How can I bid on the website?',
    'How many contract source files does the project contain?',
    'How many contracts are deployed on Arbitrum for Cosmic Signature?',
    'What is the Cosmic Signature project?',
    'How does prize distribution work?',
  ]

  const handleSuggestedQuestion = (question: string) => {
    setInputValue(question)
  }

  return (
    <div className={styles.chatContainer}>
      <div className={styles.messagesContainer}>
        {messages.map((message) => (
          <div
            key={message.id}
            className={`${styles.message} ${
              message.sender === 'user' ? styles.userMessage : styles.botMessage
            }`}
          >
            <div className={styles.messageContent}>
              {message.sender === 'bot' ? (
                <MarkdownRenderer content={message.text} />
              ) : (
                <p>{message.text}</p>
              )}
              {message.sources && message.sources.length > 0 && (
                <details className={styles.sources}>
                  <summary>Sources ({message.sources.length})</summary>
                  <ul>
                    {message.sources.map((source, idx) => (
                      <li key={idx}>{source}</li>
                    ))}
                  </ul>
                </details>
              )}
            </div>
            <span className={styles.timestamp}>
              {message.timestamp.toLocaleTimeString()}
            </span>
          </div>
        ))}
        {isLoading && (
          <div className={`${styles.message} ${styles.botMessage}`}>
            <div className={styles.messageContent}>
              <div className={styles.loader}>
                <span></span>
                <span></span>
                <span></span>
              </div>
            </div>
          </div>
        )}
        <div ref={messagesEndRef} />
      </div>

      {messages.length === 1 && (
        <div className={styles.suggestedQuestions}>
          <p className={styles.suggestionsTitle}>Try asking:</p>
          <div className={styles.suggestionButtons}>
            {suggestedQuestions.map((question, idx) => (
              <button
                key={idx}
                onClick={() => handleSuggestedQuestion(question)}
                className={styles.suggestionButton}
              >
                {question}
              </button>
            ))}
          </div>
        </div>
      )}

      <form onSubmit={handleSubmit} className={styles.inputContainer}>
        <input
          type="text"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          placeholder="Ask a question about Cosmic Signature..."
          className={styles.input}
          disabled={isLoading}
        />
        <button
          type="submit"
          className={styles.sendButton}
          disabled={!inputValue.trim() || isLoading}
        >
          Send
        </button>
      </form>
    </div>
  )
}
