import { ChakraProvider } from '@chakra-ui/react'
import { theme } from './theme'
import './App.css'

function App() {
  return (
    <ChakraProvider theme={theme}>
     <div>sample</div>
    </ChakraProvider>
  )
}

export default App
