var gulp = require('gulp');
var eslint = require('gulp-eslint');

gulp.task('eslint', function() {
  var files = [
    'public/**/*.js',
    'gulp/**/*.js',
    'gulpfile.js'
  ];

  return gulp.src(files)
    .pipe(eslint())
    .pipe(eslint.failOnError());
});
