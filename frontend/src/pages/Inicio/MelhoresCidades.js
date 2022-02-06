import { Box, chakra, Flex, Heading } from "@chakra-ui/react";
import Rating from "../../Components/Rating";

export default function MelhoresCidades() {
  return (
    <Box>
      <Heading as={'h3'} fontSize={'22'}>Melhores cidades</Heading>
      <chakra.p mb={3} color={'gray.600'}>🌆 Navegue pelas ultimas cidades que teve as melhores avaliações.</chakra.p>

      <Flex maxW="md" mx="auto" my={4}>
        <Box>
          <chakra.h3 fontSize="1xl" fontWeight="semibold" color={"gray.800"}>Juazeiro do Norte, CE</chakra.h3>
          <Rating rating={4} numReviews={4} />
        </Box>
      </Flex>

      <Flex maxW="md" mx="auto" my={4}>
        <Box>
          <chakra.h3 fontSize="1xl" fontWeight="semibold" color={"gray.800"}>Mombaça, CE</chakra.h3>
          <Rating rating={3} numReviews={5} />
        </Box>
      </Flex>

      <Flex maxW="md" mx="auto" my={4}>
        <Box>
          <chakra.h3 fontSize="1xl" fontWeight="semibold" color={"gray.800"}>Quixada, CE</chakra.h3>
          <Rating rating={3.5} numReviews={8} />
        </Box>
      </Flex>

      <Flex maxW="md" mx="auto" my={4}>
        <Box>
          <chakra.h3 fontSize="1xl" fontWeight="semibold" color={"gray.800"}>Crato, CE</chakra.h3>
          <Rating rating={4} numReviews={36} />
        </Box>
      </Flex>
    </Box>
  );
}