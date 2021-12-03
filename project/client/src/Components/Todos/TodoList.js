import { UnorderedList, ListItem } from "@chakra-ui/react"
import axios from "axios"
import { useQuery } from "react-query"

const TodoList = () => {
    const fetchTodos = async () => {
        const { data } = await axios.get('/api/todos')    
        return data
      }
    
    const { data } = useQuery("todos", fetchTodos)

    return (
        <UnorderedList>
            {data?.map(todo => <ListItem>{todo}</ListItem>)}
        </UnorderedList>
    )
}

export default TodoList
