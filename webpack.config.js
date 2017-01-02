/**
 * Created by charleen on 12/27/16.
 */
// const path= require ('path');

module.exports = {
  // This is the "main" file which should include all other modules
  entry: './client/main.js',
  // Where should the compiled file go?
  output: {
    // To the `dist` folder
    path: './client/dist',
    // With the filename `build.js` so it's dist/build.js
    filename: 'bundle.js'
  },
  resolve: {
    // NPM by default installs Runtime Only version, which will not compile html templates
    alias: {
      // this is the solution.
      'vue$': 'vue/dist/vue.common.js'
    }
  },
  module: {
    // Special compilation rules
    loaders: [
      {
        // Ask webpack to check: If this file ends with .js, then apply some transforms
        test: /\.js$/,
        // Transform it with babel
        loader: 'babel',
        // don't transform node_modules folder (which don't need to be compiled)
        exclude: /node_modules/
      },
      {
        test: /\.vue$/,
        loader: 'vue'
      }
    ]
  },
  vue: {
    loaders: {
      js: 'babel'
    }
  }
};
