import { useRef } from "react"

/**
 * Get an item from the local storage on component's first render.
 * The item will be removed from the local storage immediatly,
 * and its value will be returned and kept unchanged.
 */
export function useLocalStorageOnce(name: string) {
  const triggeredRef = useRef(false)
  const valueRef = useRef<string | null>(null)
  if (!triggeredRef.current) {
    triggeredRef.current = true
    const value = localStorage.getItem(name)
    if (value !== null) {
      valueRef.current = localStorage.getItem(name)
      localStorage.removeItem(name)
    }
  }

  return valueRef.current
}
