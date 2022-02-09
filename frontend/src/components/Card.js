import { Box, chakra, Heading, Tag } from "@chakra-ui/react";
import Rating from "./Rating";

export default function Card() {
  return (
    <Box 
      backgroundColor={'whiteAlpha.600'}
      boxShadow={'sm'}
      my='3'
      pb='3'
      borderBottom='2px solid rgba(0,0,0,0.08)'
    >
      <Heading fontSize={23}>Otimo atendimento, mas falta qualidade no serviço.</Heading>
      <chakra.p color={'gray.700'} my='1'><i>Isaque Veras</i></chakra.p>
      <chakra.p color={'gray.700'} mb='2'>lorem ipsum dolor sit amet consectetur lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt adipiscing elit sed do eiusmod tempor incididunt lorem ipsum dolor sit amet consectetur adipiscing elit</chakra.p>
      <Rating rating={5} numReviews={20} />
      <Tag size={'md'} variant='subtle' colorScheme='blue' mt='3' mr='2'>Alimentação</Tag>
      <Tag size={'md'} variant='subtle' colorScheme='green' mt='3' mr='2'>Farmacia</Tag>
    </Box>
  );
}