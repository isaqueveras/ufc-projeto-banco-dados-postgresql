import { Alert, AlertIcon, Button, chakra, CloseButton, Container, Grid, GridItem, Heading, Input } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";

import MenuDashboard from "../../../components/MenuDashboard";
import api from "../../../services/api";

export default function EditarCidades() {
	const { id } = useParams()

	const [nome, adicionarNome] = useState()
	const [uf, adicionarUF] = useState()
	const [dadosRec, adicionarDadosRec] = useState({})
  const [mensagem, adicionarMensagem] = useState({})

  useEffect(()  => {
		const fetchData = async () => {
			// Busca os dados de uma cidade utilizando a API feito em Golang
			const res = await api.get('cidade/'+id);

			adicionarNome(res?.data.nome)
			adicionarUF(res?.data.uf)
			adicionarDadosRec(res.data)
		}
		
    fetchData().catch(console.error);
  }, [id]);

	var infoEditado = {
		nome: nome !== '' ? nome : dadosRec?.nome,
		uf: uf !== '' ? uf : dadosRec?.uf,
	}
	
  function editarEmpresa() {
    api.put(`cidade/${id}`, infoEditado).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Cidade com o id (${id}) foi editado com sucesso.`,
        estaAtivo: true,
      });
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao editar a cidade com o id (${id}).`,
        estaAtivo: true
      });
    });
  }

	return (
		<>
			<MenuDashboard />
			<Container maxW='8xl' my={8}>
				{mensagem?.estaAtivo === true && (
          <Alert status={mensagem?.tipo} variant='top-accent' my={5}>
            <AlertIcon />{mensagem?.mensagem}<CloseButton onClick={() => adicionarMensagem({estaAtivo: false})} position='absolute' right='8px'/>
          </Alert>
        )}
				
				<Heading mb={5}>Editar dados da cidade</Heading>

				<Grid templateColumns='repeat(3, 1fr)' gap={6}>
					<GridItem w='100%'>
						<chakra.b>Nome da cidade</chakra.b>
						<Input variant='outline' placeholder='Nome' value={nome} onChange={e => adicionarNome(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>UF</chakra.b>
						<Input variant='outline' placeholder='UF' value={uf} onChange={e => adicionarUF(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<Button colorScheme='blue' size='md' onClick={() => editarEmpresa()}>Editar dados</Button>{' '}
						<Link to="../dashboard/empresas"><Button colorScheme='gray' size='md'>Cancelar</Button></Link>
					</GridItem>
				</Grid>
			</Container>
		</>
	)
}