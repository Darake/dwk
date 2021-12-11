import { UnorderedList } from "@chakra-ui/react"
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
        <UnorderedList>
            {data?.map(todo => <TodoRow todo={todo} />)}
        </UnorderedList>
    )
}

export default TodoList
