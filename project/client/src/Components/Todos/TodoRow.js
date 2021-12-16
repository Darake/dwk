import { CloseButton, HStack, Text } from "@chakra-ui/react"
import axios from "axios"
import { useMutation, useQueryClient } from "react-query"

const TodoRow = ({ todo }) => {
    const queryClient = useQueryClient()
      
    const todoMutation = useMutation(() => axios.put(`api/todos/${todo.id}`))

    const handleTodoCompletion = () => {
        todoMutation.mutate('todos', {
            onSuccess: () => {
                queryClient.invalidateQueries(['todos'])
            }
        })
    }

    return (
        <HStack align="flex-start">
            <CloseButton size='sm' onClick={handleTodoCompletion}/>
            <Text>{todo.description}</Text>
        </HStack>
    )
}

export default TodoRow
