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
    // Busca uma lista de cidades utilizando a API feito em Golang
    const res = await api.get('cidades');
    adicionarDados(res.data)
  }
  
  useEffect(()  => {
    fetchData().catch(console.error);
  }, []);

  function excluirCidade(id) {
    api.delete('cidade/'+ id).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Cidade com o id (${id}) foi excluido com sucesso.`,
        estaAtivo: true,
      });

      fetchData().catch(console.error);
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao excluir a cidade com o id (${id}).`,
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
          <Heading mb={5}>({dados?.total_cidades}) Cidades</Heading>
          <Link to="../dashboard/cidades/cadastrar"><Button colorScheme='green' size='md' my='1'>Cadastrar cidade</Button></Link>
        </Box>
        <Table variant='striped' size={'sm'}>
          <Thead>
            <Tr>
              <Th>ID</Th>
              <Th>Nome</Th>
              <Th>UF</Th>
              <Th>Ações</Th>
            </Tr>
          </Thead>
          <Tbody>
          {dados?.dados?.map((i) => (
            <Tr key={i.id}>
              <Td>{i?.id}</Td>
              <Td>{i?.nome}</Td>
              <Td>{i?.uf}</Td>
              <Td>{i?.data_criacao}</Td>
              <Td>
                <Link to={`../dashboard/cidades/editar/${i.id}`}><Button colorScheme='yellow' size='xs' my='1'>Editar</Button></Link>{' '}
                <Button colorScheme='red' size='xs' onClick={() => excluirCidade(i.id)} my='1'>Excluir</Button>
              </Td>
            </Tr>
            ))}
          </Tbody>
        </Table>
      </Container>
    </div>
  )
}