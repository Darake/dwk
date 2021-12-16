import { VStack } from "@chakra-ui/react"
import axios from "axios"
import { useQuery } from "react-query"
import TodoRow from "./TodoRow"

const TodoList = () => {
    const fetchTodos = async () => {
        const { data } = await axios.get('/api/todos')    
        return data
      }
    
    const { data } = useQuery("todos", fetchTodos)

    return (
        <VStack align="flex-start">
            {data?.map(todo => <TodoRow key={todo.id} todo={todo} />)}
        </VStack>
    )
}

export default TodoList
