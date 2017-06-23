const path = require('path')
const ExtractTextPlugin = require('extract-text-webpack-plugin')
const CopyWebpackPlugin = require('copy-webpack-plugin')

const extractSass = new ExtractTextPlugin({
  filename: '[name].css',
  disable: false,
  allChunks: true,
})

const copyImages = new CopyWebpackPlugin([
  {context: 'client/img', from: '**/*', to: 'img/'},
])

module.exports = {
  entry: {
    bundle: path.resolve(__dirname, 'client'),
    style: path.resolve(__dirname, 'client', 'style', 'main.scss'),
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: '[name].js',
  },
  resolve: {
    extensions: [ '.js', '.jsx', '.json' ],
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        include: [
          path.resolve(__dirname, 'app'),
          path.resolve(__dirname, 'client'),
        ],
        loader: 'babel-loader',
        query: {
          presets: [ 'es2015', 'react', 'stage-0' ],
        },
      },
      {
        test: /\.s[ac]ss$/,
        include: [
          path.resolve(__dirname, 'app'),
          path.resolve(__dirname, 'client'),
        ],
        loader: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: 'css-loader!sass-loader',
        }),
      },
    ],
  },
  plugins: [
    extractSass,
    copyImages,
  ],
}
