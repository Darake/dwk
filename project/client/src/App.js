import { VStack } from '@chakra-ui/layout';
import { QueryClientProvider, QueryClient } from 'react-query'
import DailyImage from './Components/DailyImage';
import Todos from './Components/Todos';

function App() {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <div className="App">
        <header className="App-header">
          <VStack>
            <DailyImage />
            <Todos />
          </VStack>
        </header>
      </div>
    </QueryClientProvider>
  );
}

export default App;
