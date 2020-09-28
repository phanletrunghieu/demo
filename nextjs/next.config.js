const withPlugins = require('next-compose-plugins');
const withSass = require('@zeit/next-sass')
const withLess = require('@zeit/next-less')
const lessToJS = require('less-vars-to-js')
const fs = require('fs')
const path = require('path')

const themeVariables = lessToJS(
    fs.readFileSync(path.resolve(__dirname, './styles/antd-custom.less'), 'utf8')
)

const nextConfig = {
    webpack: (config, { isServer }) => {
        if (isServer) {
            const antStyles = /antd\/.*?\/style.*?/
            const origExternals = [...config.externals];
            config.externals = [
                (context, request, callback) => {
                    if (request.match(antStyles)) return callback();
                    if (typeof origExternals[0] === 'function') {
                        origExternals[0](context, request, callback);
                    } else {
                        callback();
                    }
                },
                ...(typeof origExternals[0] === 'function' ? [] : origExternals),
            ];
    
            config.module.rules.unshift({
                test: antStyles,
                use: 'null-loader',
            });
        }
        return config;
    },
};

module.exports = withPlugins(
    [
        [
            withLess,
            {
                lessLoaderOptions: {
                    javascriptEnabled: true,
                    modifyVars: themeVariables, // make your antd custom effective
                },
            }
        ],
        [
            withSass,
            {
                cssModules: true,
                cssLoaderOptions: {
                    localIdentName: '[path]___[local]___[hash:base64:5]',
                },
            },
        ],
    ],
    nextConfig,
);