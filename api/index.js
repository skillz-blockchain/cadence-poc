import app from 'cadence-web/server/index.js';
import router from 'cadence-web/server/routes.js';
import { initRouter/*, initMiddleware*/ } from './server';

const port = process.env.PORT;

initRouter(router);
//initMiddleware(app);

app.init({ useWebpack: false }).listen(port);
console.log('node server up and listening on port ' + port);
