const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: './dashboard/index.js',
  output: {
    path: path.resolve(__dirname, 'public'),
    filename: 'assets/bundle.js',
    publicPath: '/',
  },
  plugins: [new HtmlWebpackPlugin({
    template: './dashboard/index.html',
    inject: 'body',
  })],
  module: {
    loaders: [
      { test: /\.js$/, loader: 'babel-loader', exclude: /node_modules/ },
      { test: /\.jsx$/, loader: 'babel-loader', exclude: /node_modules/ },
      { test: /\.css$/, use: ['style-loader', 'css-loader'] },
    ],
  },
};
