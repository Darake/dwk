import { Input, Button, FormControl, FormErrorMessage } from "@chakra-ui/react"
import axios from "axios"
import { useForm } from "react-hook-form";
import { useMutation, useQueryClient } from "react-query";

const TodoCreation = () => {
    const {
        handleSubmit,
        register,
        formState: { errors }
    } = useForm();

    const queryClient = useQueryClient()
      
    const todoMutation = useMutation((data) => axios.post('api/todos', data))

    const handleTodoCreation = values => {
        todoMutation.mutate(values.todo, {
            onSuccess: () => {
                queryClient.invalidateQueries(['todos'])
            }
        })
    }

    return (
        <form onSubmit={handleSubmit(handleTodoCreation)}>
            <FormControl isInvalid={errors.todo}>
                <Input
                    id="todo"
                    {...register("todo", {
                        required: "This is required",
                        maxLength: { value: 160, message: "Maximum length should be 160" }
                    })}
                />
                <FormErrorMessage>
                    {errors.todo && errors.todo.message}
                </FormErrorMessage>
            </FormControl>
            <Button mt={4} colorScheme="teal" type="submit">
                Create todo
            </Button>
        </form>
    );

}

export default TodoCreation
