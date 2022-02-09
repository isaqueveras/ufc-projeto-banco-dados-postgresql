import { useEffect, useState } from "react";
import { 
  Alert,
  AlertIcon,
  Button,
  CloseButton,
  Container,
  Heading,
  Table, 
  Tbody, 
  Td, 
  Th, 
  Thead, 
  Tr,
} from "@chakra-ui/react";

// Componentes
import MenuDashboard from "../../../components/MenuDashboard";

// Serviços
import api from "../../../services/api";
import { Link } from "react-router-dom";

export default function Inicio() {
  const [dados, adicionarDados] = useState([]);
  const [mensagem, adicionarMensagem] = useState({})

  const fetchData = async () => {
    // Busca uma lista de empresa utilizando a API feito em Golang
    const res = await api.get('empresas');
    adicionarDados(res.data)
  }
  
  useEffect(()  => {
    fetchData().catch(console.error);
  }, []);

  function excluirEmpresa(id) {
    api.delete('empresa/'+ id).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Empresa com o id (${id}) foi excluido com sucesso.`,
        estaAtivo: true,
      });

      fetchData().catch(console.error);
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao excluir a empresa com o id (${id}).`,
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

        <Heading mb={5}>({dados?.total_empresas}) Empresas</Heading>
        <Table variant='striped'>
          <Thead>
            <Tr>
              <Th>ID</Th>
              <Th>Nome</Th>
              <Th>Telefone</Th>
              <Th>CPF</Th>
              <Th>CNPJ</Th>
              <Th>ID Cidade</Th>
              <Th>ID Categoria</Th>
              <Th>Data criação</Th>
              <Th>Ações</Th>
            </Tr>
          </Thead>
          <Tbody>
          {dados?.dados?.map((i) => (
            <Tr key={i.id}>
              <Td>{i?.id}</Td>
              <Td>{i?.nome}</Td>
              <Td>{i?.telefone}</Td>
              <Td>{i?.cpf}</Td>
              <Td>{i?.cnpj}</Td>
              <Td>{i?.cidade_id}</Td>
              <Td>{i?.categoria_id}</Td>
              <Td>{i?.data_criacao}</Td>
              <Td>
                <Link to={`editar/${i.id}`}><Button colorScheme='yellow' size='xs' my='1'>Editar</Button></Link>{' '}
                <Button colorScheme='red' size='xs' onClick={() => excluirEmpresa(i.id)} my='1'>Excluir</Button>
              </Td>
            </Tr>
            ))}
          </Tbody>
        </Table>
      </Container>
    </div>
  )
}