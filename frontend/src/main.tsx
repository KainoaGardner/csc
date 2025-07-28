import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import { AppProvider } from "./appContext/appProvider.tsx"
import { SettingProvider } from "./settingContext/settingProvider.tsx"

import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AppProvider>
      <SettingProvider>
        <App />
      </SettingProvider>
    </AppProvider>
  </StrictMode>,
)
