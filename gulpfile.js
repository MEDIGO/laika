'use strict';

var gulp = require('gulp');

require('gulp-modularize')('./gulp/');

gulp.task('default', ['validate']);
