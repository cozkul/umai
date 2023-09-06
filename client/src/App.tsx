import { Box, List, ThemeIcon } from '@mantine/core';
import useSWR from "swr";
import './App.css';
import AddFact from "./components/AddFact";
import { CheckCircleFillIcon } from '@primer/octicons-react';

export interface Fact {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date;
  question: string;
  answer: string;
}

export const ENDPOINT = 'http://localhost:3000'

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {
  // SWR is a react hook for data fetching
  const { data, mutate } = useSWR<Fact[]>(" ", fetcher)

  // Mantine define style using sx
  return (
    <Box
      sx={(theme) => ({
        padding: "2rem",
        width: "100%",
        maxWidth: "40rem",
        margin: "0 auto",
      })}
    >
      <List spacing="xs" size="sm" mb={12} center>
        {data?.map((fact: Fact) => {
          // prefixing keys allows them to be unique
          return <List.Item key={`fact__${fact.ID}`} icon={
              fact.ID % 2 == 0 ? (<ThemeIcon color='teal' size={24} radius="xl">
                <CheckCircleFillIcon size={20} />
              </ThemeIcon>) : (<ThemeIcon color='gray' size={24} radius="xl">
                <CheckCircleFillIcon size={20} />
              </ThemeIcon>)
            }>
            {fact.question}
          </List.Item>
        })}
      </List>
      <AddFact mutate={mutate} />
    </Box>
  );

}

export default App;
