import './globals.css'
import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Cosmic Signature FAQ Bot',
  description: 'AI-powered FAQ bot for Cosmic Signature project',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
