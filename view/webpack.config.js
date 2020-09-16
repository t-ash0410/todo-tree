const path = require('path');
const VueLoaderPlugin = require("vue-loader/lib/plugin");

module.exports = {
  mode: "development",
  entry: ["babel-polyfill", path.resolve("src", "script", "index.ts")],
  output: {
    filename: "bundle.js",
    path: path.join(__dirname, "static/script/")
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: "vue-loader"
      },
      {
        test: /\.js$/,
        loader: "babel-loader"
      },
      {
        test: /\.ts$/,
        use: [{
            loader: 'ts-loader',
            options: {
                appendTsSuffixTo: [/\.vue$/]
            }
        }]
      },
      {
        test: /\.s[ac]ss$/i,
        use: [
          'vue-style-loader',
          'css-loader',
          'sass-loader',
        ],
      }
    ]
  },
  resolve: {
    extensions: [".js", ".ts", "json", "jsx", "vue"],
    alias: {
      vue$: "vue/dist/vue.esm.js"
    }
  },
  devServer: {
    contentBase: "static",
    proxy: {
      "/api": "http://localhost:9080"
    }
  },
  plugins: [
    new VueLoaderPlugin()
  ]
};
