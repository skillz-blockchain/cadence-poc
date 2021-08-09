import {
    deployValidators,
    handleError,
} from '../cadence/...';

export async function deploy(ctx) {
    const { cadence, request } = ctx;
    const { description, name } = request.body;

    try {
        const product = await deployValidators({
            cadence,
            description,
            name,
        });
        ctx.body = product;
    } catch (error) {
        return handleError({ ctx, error });
    }
}

export async function unstake(ctx) {
}