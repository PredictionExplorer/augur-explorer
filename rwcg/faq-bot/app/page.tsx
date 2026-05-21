'use client'

import { useState } from 'react'
import ChatInterface from './components/ChatInterface'
import styles from './page.module.css'

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.container}>
        <header className={styles.header}>
          <h1>Cosmic Signature FAQ Bot</h1>
          <p>Ask me anything about the Cosmic Signature project</p>
        </header>
        <ChatInterface />
      </div>
    </main>
  )
}
