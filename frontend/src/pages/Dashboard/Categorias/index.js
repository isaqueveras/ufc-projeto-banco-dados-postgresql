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
    const res = await api.get('categorias');
    adicionarDados(res.data)
  }
  
  useEffect(()  => {
    fetchData().catch(console.error);
  }, []);

  function excluirCategoria(id) {
    api.delete('categoria/'+ id).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Categoria com o id (${id}) foi excluido com sucesso.`,
        estaAtivo: true,
      });

      fetchData().catch(console.error);
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao excluir a categoria com o id (${id}).`,
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
          <Heading mb={5}>({dados?.total_categorias}) Categorias</Heading>
          <Link to="../dashboard/categorias/cadastrar"><Button colorScheme='green' size='md' my='1'>Cadastrar categoria</Button></Link>
        </Box>
        <Table variant='striped' size={'sm'}>
          <Thead>
            <Tr>
              <Th>ID</Th>
              <Th>Nome</Th>
              <Th>Descrição</Th>
              <Th>Ações</Th>
            </Tr>
          </Thead>
          <Tbody>
          {dados?.dados?.map((i) => (
            <Tr key={i.id}>
              <Td>{i?.id}</Td>
              <Td>{i?.nome}</Td>
              <Td>{i?.descricao}</Td>
              <Td>{i?.data_criacao}</Td>
              <Td>
                <Link to={`../dashboard/categorias/editar/${i.id}`}><Button colorScheme='yellow' size='xs' my='1'>Editar</Button></Link>{' '}
                <Button colorScheme='red' size='xs' onClick={() => excluirCategoria(i.id)} my='1'>Excluir</Button>
              </Td>
            </Tr>
            ))}
          </Tbody>
        </Table>
      </Container>
    </div>
  )
}