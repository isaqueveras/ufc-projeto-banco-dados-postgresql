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
  Flex, 
} from "@chakra-ui/react";

function Home() {
  return (
    <Container maxW='8xl' height={'100vh'}>
      <Grid templateColumns='repeat(1, 1fr)' gap={0}>
        <GridItem backgroundColor="white" maxW={'600px'} width={600}>
          <Box py={8}>
            <chakra.p mb={2} fontSize="xs" fontWeight="semibold" letterSpacing="wide" color="gray.400" textTransform="uppercase">Para consumidores</chakra.p>
            <chakra.h1 mb={3} fontSize={{ base: "3xl", md: "4xl" }} fontWeight="bold" lineHeight="shorter" color={useColorModeValue("gray.900", "white")}>Encontre todas as avaliações</chakra.h1>
            <chakra.p mb={5} color="gray.500" fontSize={{ md: "lg" }}>
              Today every company needs apps to engage their customers and run their
              businesses. Step up your ability to build, manage, and deploy great
            </chakra.p>
            <HStack>
              <Button
                as="a"
                w={{ base: "full", sm: "auto" }}
                variant="solid"
                colorScheme="messenger"
                size="lg"
                mb={{ base: 2, sm: 0 }}
                cursor="pointer"
              >
                Adicionar uma avaliações
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
      
      <Grid templateColumns='repeat(3, 1fr)' gap={4}>
        <GridItem 
          backgroundColor="gray.100" 
          p={3}>
          <Heading as={'h3'} fontSize={'22'}>Cidades recentes</Heading>
          <chakra.p>Navegue pelas ultimas cidades que teve avaliações</chakra.p>

          <Flex maxW="md" mx="auto" backgroundColor={'white'} rounded="lg" mt={2}>
            <Box w={2 / 3} p={{ base: 4, md: 4 }}>
              <chakra.h3 fontSize="1xl" fontWeight="bold" color={useColorModeValue("gray.800", "white")}>Backpack</chakra.h3>
              <chakra.p fontSize="sm" color={"gray.600"}>Lorem ipsum dolor sit amet consectetur adipisicing elit In odit</chakra.p>
              <chakra.p>3 estrelas</chakra.p>
            </Box>
          </Flex>

          <Flex maxW="md" mx="auto" backgroundColor={'white'} rounded="lg" mt={2}>
            <Box w={2 / 3} p={{ base: 4, md: 4 }}>
              <chakra.h3 fontSize="1xl" fontWeight="bold" color={useColorModeValue("gray.800", "white")}>Backpack</chakra.h3>
              <chakra.p fontSize="sm" color={"gray.600"}>Lorem ipsum dolor sit amet consectetur adipisicing elit In odit</chakra.p>
              <chakra.p>3 estrelas</chakra.p>
            </Box>
          </Flex>

          <Flex maxW="md" mx="auto" backgroundColor={'white'} rounded="lg" mt={2}>
            <Box w={2 / 3} p={{ base: 4, md: 4 }}>
              <chakra.h3 fontSize="1xl" fontWeight="bold" color={useColorModeValue("gray.800", "white")}>Backpack</chakra.h3>
              <chakra.p fontSize="sm" color={"gray.600"}>Lorem ipsum dolor sit amet consectetur adipisicing elit In odit</chakra.p>
              <chakra.p>3 estrelas</chakra.p>
            </Box>
          </Flex>

          <Flex maxW="md" mx="auto" backgroundColor={'white'} rounded="lg" mt={2}>
            <Box w={2 / 3} p={{ base: 4, md: 4 }}>
              <chakra.h3 fontSize="1xl" fontWeight="bold" color={useColorModeValue("gray.800", "white")}>Backpack</chakra.h3>
              <chakra.p fontSize="sm" color={"gray.600"}>Lorem ipsum dolor sit amet consectetur adipisicing elit In odit</chakra.p>
              <chakra.p>3 estrelas</chakra.p>
            </Box>
          </Flex>

          <Flex maxW="md" mx="auto" backgroundColor={'white'} rounded="lg" mt={2}>
            <Box w={2 / 3} p={{ base: 4, md: 4 }}>
              <chakra.h3 fontSize="1xl" fontWeight="bold" color={useColorModeValue("gray.800", "white")}>Backpack</chakra.h3>
              <chakra.p fontSize="sm" color={"gray.600"}>Lorem ipsum dolor sit amet consectetur adipisicing elit In odit</chakra.p>
              <chakra.p>3 estrelas</chakra.p>
            </Box>
          </Flex>

        </GridItem>
        <GridItem 
          backgroundColor="gray.100" 
          height={'500px'}
          p={3}>
          <Heading as={'h3'} fontSize={'22'}>Avaliações</Heading>
          <chakra.p>Leia as ultimas avaliações</chakra.p>
        </GridItem>
        <GridItem 
          backgroundColor="gray.100" 
          height={'500px'}
          p={3}>
          <Heading as={'h3'} fontSize={'22'}>Categorias</Heading>
          <chakra.p>Navegue pelas categorias</chakra.p>
        </GridItem>
      </Grid>
    </Container>
  );  
}

export default Home;