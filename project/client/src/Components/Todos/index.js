import { VStack } from "@chakra-ui/react"
import TodoCreation from "./TodoCreation"
import TodoList from "./TodoList"

const Todos = () => {
  return (
    <VStack>
      <TodoCreation />
      <TodoList />
    </VStack>
  )
}

export default Todos
