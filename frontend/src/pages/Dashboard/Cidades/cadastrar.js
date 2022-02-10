import { Alert, AlertIcon, Button, chakra, CloseButton, Container, Grid, GridItem, Heading, Input } from "@chakra-ui/react";
import { useState } from "react";
import { Link } from "react-router-dom";

import MenuDashboard from "../../../components/MenuDashboard";
import api from "../../../services/api";

export default function CadastrarEmpresa() {
  const [mensagem, adicionarMensagem] = useState({})
	
	const [nome, adicionarNome] = useState()
	const [uf, adicionarUF] = useState()

	var dadosCadastrar = {
		nome: nome,
		uf: uf,
	}
	
  function cadastrarCidade() {
    api.post(`cidades`, dadosCadastrar).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Cidade cadastrada com sucesso.`,
        estaAtivo: true,
      });

			adicionarNome(''); adicionarUF('');
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao cadastrar a cidade.`,
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
				
				<Heading mb={5}>Cadastrar uma nova cidade</Heading>

				<Grid templateColumns='repeat(2, 1fr)' gap={6}>
					<GridItem w='100%'>
						<chakra.b>Nome</chakra.b>
						<Input variant='outline' placeholder='Nome' value={nome} onChange={e => adicionarNome(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>UF</chakra.b>
						<Input variant='outline' placeholder='UF' value={uf} onChange={e => adicionarUF(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						{nome?.length > 5 && uf !== '' ? (
							<Button colorScheme='blue' size='md' onClick={() => cadastrarCidade()}>Cadastrar Cidade</Button>
						) : (
							<Button colorScheme='gray' size='md'>Cadastrar Cidade</Button>
						)}
						{' '}<Link to="../dashboard/cidades"><Button colorScheme='gray' size='md'>Cancelar</Button></Link>
					</GridItem>
				</Grid>
			</Container>
		</>
	)
}