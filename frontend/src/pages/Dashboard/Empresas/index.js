import { useEffect, useState } from "react";
import { 
  Table, 
  TableCaption, 
  Tbody, 
  Td, 
  Tfoot, 
  Th, 
  Thead, 
  Tr,
} from "@chakra-ui/react";

// Componentes
import MenuDashboard from "../../../components/MenuDashboard";

// Serviços
import api from "../../../services/api";

export default function Inicio() {
  const [dados, adicionarDados] = useState([]);
  
  useEffect(()  => {
    const fetchData = async () => {
      // Busca uma lista de empresa utilizando a API feito em Golang
      const res = await api.get('empresas');
      adicionarDados(res.data)
    }
  
    fetchData().catch(console.error);
  }, []);

  return (
    <div>
      <MenuDashboard />
      <Table variant='simple'>
        <TableCaption>Imperial to metric conversion factors</TableCaption>
        <Thead>
          <Tr>
            <Th>To convert</Th>
            <Th>into</Th>
            <Th isNumeric>multiply by</Th>
          </Tr>
        </Thead>
        <Tbody>
          <Tr>
            <Td>{dados[0]?.nome ? 'tem dados' : 'Nao tem'}</Td>
            <Td>millimetres (mm)</Td>
            <Td isNumeric>25.4</Td>
          </Tr>
          <Tr>
            <Td>feet</Td>
            <Td>centimetres (cm)</Td>
            <Td isNumeric>30.48</Td>
          </Tr>
          <Tr>
            <Td>yards</Td>
            <Td>metres (m)</Td>
            <Td isNumeric>0.91444</Td>
          </Tr>
        </Tbody>
        <Tfoot>
          <Tr>
            <Th>To convert</Th>
            <Th>into</Th>
            <Th isNumeric>multiply by</Th>
          </Tr>
        </Tfoot>
      </Table>
    </div>
  )
}