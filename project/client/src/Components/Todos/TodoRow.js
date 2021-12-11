import { ListItem, CloseButton } from "@chakra-ui/react"
import axios from "axios"
import { useMutation, useQueryClient } from "react-query"

const TodoRow = ({ todo }) => {
    const queryClient = useQueryClient()
      
    const todoMutation = useMutation(() => axios.put(`api/todos/${todo.id}`))

    const handleTodoCompletion = () => {
        todoMutation.mutate({
            onSuccess: () => {
                queryClient.invalidateQueries(['todos'])
            }
        })
    }

    return (
        <ListItem>
            <CloseButton size='sm' onClick={handleTodoCompletion}/>
            {todo.description}
        </ListItem>
    )
}

export default TodoRow
