const gulp    = require("gulp")
const babel   = require("gulp-babel")
const include = require("gulp-include")
const sass    = require("gulp-sass")
const rev     = require("gulp-rev")
const shell   = require("gulp-shell")
const minify  = require("gulp-minify")
const uglify  = require("gulp-uglify")
const gulpif  = require("gulp-if")

gulp.task("assets:compile:js", function() {
  gulp.src("assets/javascripts/**/*.js")
    .pipe(include())
    .pipe(babel({presets: ["es2015"]}))
    .pipe(gulp.dest("public/assets"))
})

gulp.task("assets:compile:scss", function() {
  gulp.src("assets/stylesheets/**/*.scss")
    .pipe(include())
    .pipe(sass())
    .pipe(gulp.dest("public/assets"))
})

gulp.task("server", ["assets:compile:js", "assets:compile:scss"], function() {
  gulp.src("").pipe(shell("gin main.go"));
  gulp.watch(["app/assets/**/*.js", "app/assets/**/*.scss"], ["assets:compile"]);
});

gulp.task("test", ["assets:compile"], function() {
  gulp.src("").pipe(shell("ginkgo -r"));
});

gulp.task("assets:precompile", ["assets:precompile:js", "assets:precompile:scss"], function() {
  gulp.src("").pipe(shell("echo assets compiled"))
})

gulp.task("assets:precompile:js", function() {
  gulp.src("assets/javascripts/**/*.js")
    .pipe(include())
    .pipe(babel({presets: ["es2015"]}))
    .pipe(uglify())
    // .pipe(rev())
    // .pipe(gulp.dest("public/assets"))
    // .pipe(rev.manifest({ path: "manifest.json"}))
    .pipe(gulp.dest("public/assets"))
})

gulp.task("assets:precompile:scss", function() {
  gulp.src("assets/stylesheets/**/*.scss")
    .pipe(include())
    .pipe(sass())
    .pipe(minify())
    // .pipe(rev())
    // .pipe(gulp.dest("public/assets"))
    // .pipe(rev.manifest({ path: "manifest.json" }))
    .pipe(gulp.dest("public/assets"))
})

// gulp.task("assets:precompile", function() {
//   gulp.src("assets/**/*")
//     .pipe(include())
//     .pipe(gulpif("*.js", babel({ presets: ["es2015"] })))
//     .pipe(gulpif("*.js", uglify()))
//     .pipe(gulpif("*.scss", sass()))
//     .pipe(gulpif("*.css", minify()))
//     .pipe(rev())
//     .pipe(gulp.dest("public/assets")) // file with digest
//     .pipe(rev.manifest({path: 'manifest.json'}))
//     .pipe(gulp.dest("public/assets")) // manifest.json
// })