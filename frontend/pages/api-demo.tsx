import { useEffect, useState } from 'react'

export default function ApiDemo() {
  const [message, setMessage] = useState('Loading...')

  useEffect(() => {
    fetch('/api/hello')
      .then((res) => res.json())
      .then((data) => setMessage(data.message))
      .catch(() => setMessage('Error connecting to API'))
  }, [])

  return (
    <main style={{ padding: 40 }}>
      <h1>Harmonia Dashboard</h1>
      <p>Backend says: {message}</p>
    </main>
  )
}
