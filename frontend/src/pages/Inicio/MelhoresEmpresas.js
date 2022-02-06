import { Box, chakra, Flex, Heading } from "@chakra-ui/react";
import Rating from "../../Components/Rating";

export default function MelhoresEmpresas() {
  return (
    <Box mt={8}>
      <Heading as={'h3'} fontSize={'22'}>Melhores empresas</Heading>
      <chakra.p mb={3} color={'gray.600'}>👛 Conheça as empresas que tem as melhores avaliações</chakra.p>
      <Flex maxW="md" mx="auto" my={4}>
        <Box>
          <chakra.h3 fontSize="1xl" fontWeight="semibold" color={"gray.800"}>Brisanet Telecomunicações</chakra.h3>
          <Rating rating={5} numReviews={365} />
        </Box>
      </Flex>
      <Flex maxW="md" mx="auto" my={4}>
        <Box>
          <chakra.h3 fontSize="1xl" fontWeight="semibold" color={"gray.800"}>Coca-Cola LTDA</chakra.h3>
          <Rating rating={4} numReviews={452} />
        </Box>
      </Flex>
      <Flex maxW="md" mx="auto" my={4}>
        <Box>
          <chakra.h3 fontSize="1xl" fontWeight="semibold" color={"gray.800"}>João Eletricista</chakra.h3>
          <Rating rating={3} numReviews={132} />
        </Box>
      </Flex>
    </Box>
  );
}