import { 
  Box, 
  Button, 
  chakra, 
  HStack, 
  useColorModeValue, 
  Container, 
  Grid,
  GridItem,
  Heading,
  Tag,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";

import Menu from '../../components/Menu'
import api from "../../services/api";

import MelhoresCidades from "./MelhoresCidades";
import MelhoresEmpresas from "./MelhoresEmpresas";
import UltimasAvaliacoes from "./UltimasAvaliacoes";
import PioresEmpresas from "./PioresEmpresas";

function Home() {
  const [dadosCategorias, adicionarDadosCategorias] = useState([])

  useEffect(()  => {
    const fetchCategorias = async () => {
      const res = await api.get('categorias');
      adicionarDadosCategorias(res.data)
    }

    fetchCategorias().catch(console.error);
  }, []);

  return (
    <Box>
      <Menu />
      <Container maxW='8xl'>
        <Grid templateColumns='repeat(1, 1fr)' gap={0}>
          <GridItem>
            <Box py={12}>
              <chakra.p mb={2} fontSize="xs" fontWeight="semibold" letterSpacing="wide" color="gray.400" textTransform="uppercase">Para consumidores</chakra.p>
              <chakra.h1 fontSize={{ base: "3xl", md: "4xl" }} fontWeight="bold" lineHeight="shorter" color={useColorModeValue("gray.900", "white")}>🎉 A avalição que você precisa está aqui!</chakra.h1>
              <chakra.p mb={2} color="gray.500" fontSize={{ md: "lg" }}>
                Somos um centro de avaliações onde os consumidores podem encontrar todo tipo de avalições. <br /> Avaliações de comerciantes, prestadores de serviços, empresas grandes, produtos e etc.
              </chakra.p>
              <HStack>
                <Button
                  as="a"
                  w={{ base: "full", sm: "auto" }}
                  variant="solid"
                  colorScheme="yellow"
                  size="lg"
                  mb={{ base: 2, sm: 0 }}
                  cursor="pointer"
                >
                  📃 Ler avaliações
                </Button>
                <Button
                  as="a"
                  w={{ base: "full", sm: "auto" }}
                  mb={{ base: 2, sm: 0 }}
                  size="lg"
                  cursor="pointer"
                  color="blackAlpha.800"
                >
                  Sobre o projeto
                </Button>
              </HStack>
            </Box>
          </GridItem>
        </Grid>
        
        <Grid mb={5} templateColumns='repeat(4, 1fr)' gap={4}>
          <GridItem rounded="lg">
            <MelhoresCidades />
            <MelhoresEmpresas />
          </GridItem>
          <GridItem rounded="lg" colSpan={2}>
            <Heading as={'h3'} fontSize={'22'}>Avaliações</Heading>
            <chakra.p mb={6} color={'gray.600'}>🛍️ Leia as ultimas avaliações</chakra.p>
            <UltimasAvaliacoes />
          </GridItem>
          <GridItem rounded="lg">
            <Heading as={'h3'} fontSize={'22'}>Categorias</Heading>
            <chakra.p mb={6} color={'gray.600'}>🏬 Navegue pelas categorias</chakra.p>

            {dadosCategorias?.dados?.map((i) => (
              <Box key={i?.dados?.id}> 
                <Tag size={'md'} variant='subtle' colorScheme='green' mt='3' mr='2'>{i?.nome}</Tag>
              </Box>
            ))}

            <PioresEmpresas />
          </GridItem>
        </Grid>
      </Container>
    </Box>
  );  
}

export default Home;