const fs = require('fs');
const content = JSON.parse(fs.readFileSync('package-lock.json', 'utf-8'));

// Skip integrity check for service-exchange-webapp, which differs on Apple Silicon M1.
// @see https://github.com/npm/cli/issues/2846
delete content.dependencies['service-exchange-webapp'].integrity;
delete content.packages['node_modules/service-exchange-webapp'].integrity;

fs.writeFileSync('package-lock.json', JSON.stringify(content, null, 2) + '\n');
