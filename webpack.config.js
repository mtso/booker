const path = require('path')

module.exports = {
  entry: {
    bundle: path.resolve(__dirname, 'client')
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
          presets: [ 'es2015', 'react' ],
        },
      },
    ],
  },
}
