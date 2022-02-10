import { useEffect, useState } from "react";
import {  Alert, AlertIcon, Box, Button, CloseButton, Container, Heading, Table,  Tbody,  Td,  Th,  Thead,  Tr,} from "@chakra-ui/react";

// Componentes
import MenuDashboard from "../../../components/MenuDashboard";

// Serviços
import api from "../../../services/api";
import { Link } from "react-router-dom";

export default function Inicio() {
  const [dados, adicionarDados] = useState([]);
  const [mensagem, adicionarMensagem] = useState({})

  const fetchData = async () => {
    // Busca uma lista de categorias utilizando a API feito em Golang
    const res = await api.get('avaliacoes');
    adicionarDados(res.data)
  }
  
  useEffect(()  => {
    fetchData().catch(console.error);
  }, []);

  function excluirAvaliacao(id) {
    api.delete('avaliacao/'+ id).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Avaliação com o id (${id}) foi excluido com sucesso.`,
        estaAtivo: true,
      });

      fetchData().catch(console.error);
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao excluir a avaliação com o id (${id}).`,
        estaAtivo: true
      });
    });
  }

  return (
    <div>
      <MenuDashboard />
      <Container maxW='8xl' my={8}>

        {mensagem?.estaAtivo === true && (
          <Alert status={mensagem?.tipo} variant='top-accent' my={5}>
            <AlertIcon />{mensagem?.mensagem}<CloseButton onClick={() => adicionarMensagem({estaAtivo: false})} position='absolute' right='8px'/>
          </Alert>
        )}

        <Box
          display={'flex'}
          flexDirection={'row'}
          justifyContent={'space-between'}
        >
          <Heading mb={5}>({dados?.total_avaliacoes}) Avaliações</Heading>
          <Link to="/dashboard/avaliacoes/cadastrar"><Button colorScheme='green' size='md' my='1'>Cadastrar avaliação</Button></Link>
        </Box>
        <Table variant='striped' size={'sm'}>
          <Thead>
            <Tr>
              <Th>ID</Th>
              <Th>Titulo</Th>
              <Th>Descrição</Th>
              <Th>Empresa</Th>
              <Th>Ações</Th>
            </Tr>
          </Thead>
          <Tbody>
          {dados?.dados?.map((i) => (
            <Tr key={i.id}>
              <Td>{i?.id}</Td>
              <Td>{i?.titulo}</Td>
              <Td>{i?.descricao}</Td>
              <Td>(#{i?.empresa?.id}) {i?.empresa?.Nome}</Td>
              <Td>{i?.data_criacao}</Td>
              <Td>
                <Link to={`/dashboard/avaliacoes/editar/${i.id}`}><Button colorScheme='yellow' size='xs' my='1'>Editar</Button></Link>{' '}
                <Button colorScheme='red' size='xs' onClick={() => excluirAvaliacao(i.id)} my='1'>Excluir</Button>
              </Td>
            </Tr>
            ))}
          </Tbody>
        </Table>
      </Container>
    </div>
  )
}