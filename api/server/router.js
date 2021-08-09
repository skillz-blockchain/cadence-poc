import { register, login } from '../controllers/users';
import { deploy, unstake } from '../controllers/validators';


// Mock up taken from https://www.notion.so/skillzblockchain/PRD-Ethereum-2-0-Validator-node-API-2c1023a26dcb4bf99e927c24596e2be6#a9c95e7c71f149c78065ed6d92b45c23
const initRouter = (router) => {
    router.post('/v1/register', register)
    router.post('/v1/login', login)

    router.post('/v1/deploy/eth2', deploy)
    router.post('/v1/unstake/eth2', unstake)
}

export default initRouter;