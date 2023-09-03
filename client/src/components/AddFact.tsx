import { useState } from 'react'
import { useForm } from '@mantine/form'
import { Button, Group, Modal, TextInput, Textarea } from '@mantine/core'
import { ENDPOINT, Fact } from '../App';
import { KeyedMutator } from 'swr';

// function AddFact({mutate}: KeyedMutator<any> BAD Practice - use interface instead) {
function AddFact({ mutate }: { mutate: KeyedMutator<Fact[]> }) {
    // Control whether addFact dialog box is open or close
    // useState is react hook
    const [open, setOpen] = useState(false)

    const form = useForm({
        initialValues: {
            question: "",
            answer: "",
        },
    })

    async function createFact(values: { question: string, answer: string }) {
        const updated = await fetch(`${ENDPOINT}/fact`, {
            method: 'POST',
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(values),
        }).then(r => r.json())

        // Calling mutate with newly updated list updates returned result of useSWR
        mutate(updated);
        form.reset();
        setOpen(false);
    }

    return (
        <>
            <Modal
                opened={open}
                onClose={() => setOpen(false)}
                title="Write down your fact man"
            >
                <form onSubmit={form.onSubmit(createFact)}>
                    <TextInput
                        required
                        mb={12}
                        label="Question"
                        placeholder="What is the question"
                        {...form.getInputProps("question")}
                    />
                    <Textarea
                        required
                        mb={12}
                        label="Fact"
                        placeholder="Write answer here"
                        {...form.getInputProps("answer")}
                    />
                    <Button type="submit">Create Fact</Button>
                </form>
            </Modal>

            <Group position="center">
                <Button
                    fullWidth
                    mb={12}
                    onClick={() => setOpen(true)}
                >
                    ADD FACT
                </Button>
            </Group>
        </>
    );
}

export default AddFact