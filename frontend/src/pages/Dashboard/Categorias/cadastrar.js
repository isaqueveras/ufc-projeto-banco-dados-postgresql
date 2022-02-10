import { Alert, AlertIcon, Button, chakra, CloseButton, Container, Grid, GridItem, Heading, Input } from "@chakra-ui/react";
import { useState } from "react";
import { Link } from "react-router-dom";

import MenuDashboard from "../../../components/MenuDashboard";
import api from "../../../services/api";

export default function CadastrarCategoria() {
  const [mensagem, adicionarMensagem] = useState({})
	
	const [nome, adicionarNome] = useState()
	const [descricao, adicionarDescricao] = useState()

	var dadosCadastrar = {
		nome: nome,
		descricao: descricao,
	}
	
  function cadastrarCidade() {
    api.post(`categorias`, dadosCadastrar).then(res => {
      adicionarMensagem({
        tipo: 'success', 
        mensagem: `Categoria cadastrada com sucesso.`,
        estaAtivo: true,
      });

			adicionarNome(''); adicionarDescricao('');
    }).catch(error => {
      adicionarMensagem({
        tipo: 'error',
        mensagem: `Erro ao cadastrar a categoria.`,
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
				
				<Heading mb={5}>Cadastrar uma nova categoria</Heading>

				<Grid templateColumns='repeat(2, 1fr)' gap={6}>
					<GridItem w='100%'>
						<chakra.b>Nome</chakra.b>
						<Input variant='outline' placeholder='Nome' value={nome} onChange={e => adicionarNome(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						<chakra.b>Descrição</chakra.b>
						<Input variant='outline' placeholder='Descrição' value={descricao} onChange={e => adicionarDescricao(e.target.value)} />
					</GridItem>
					<GridItem w='100%'>
						{nome?.length > 5 && descricao !== '' ? (
							<Button colorScheme='blue' size='md' onClick={() => cadastrarCidade()}>Cadastrar Categoria</Button>
						) : (
							<Button colorScheme='gray' size='md'>Cadastrar Categoria</Button>
						)}
						{' '}<Link to="../dashboard/cidades"><Button colorScheme='gray' size='md'>Cancelar</Button></Link>
					</GridItem>
				</Grid>
			</Container>
		</>
	)
}