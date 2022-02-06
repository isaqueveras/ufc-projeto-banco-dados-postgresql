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
} from "@chakra-ui/react";
import Card from "../../Components/Card";

import Menu from '../../Components/Menu'
import MelhoresCidades from "./MelhoresCidades";
import MelhoresEmpresas from "./MelhoresEmpresas";

function Home() {
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

            <Card />
            <Card />
            <Card />
            <Card />
            <Card />
            <Card />
            
          </GridItem>
          <GridItem   rounded="lg">
            <Heading as={'h3'} fontSize={'22'}>Categorias</Heading>
            <chakra.p mb={6} color={'gray.600'}>🏬 Navegue pelas categorias</chakra.p>
          </GridItem>
        </Grid>
      </Container>
    </Box>
  );  
}

export default Home;