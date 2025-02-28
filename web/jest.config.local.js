module.exports = {
  roots: ["<rootDir>/src"],
  transform: {
    "^.+\\.tsx?$": "ts-jest",
    "\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga|ico)$":
      "<rootDir>/file-transformer.js",
  },
  moduleNameMapper: {
    "^pipecd/(.*)$": "<rootDir>/../bazel-bin/$1",
    "^~/(.*)$": "<rootDir>/src/$1",
    "^~~/(.*)$": "<rootDir>/$1",
  },
  moduleDirectories: ["<rootDir>/node_modules", "__fixtures__"],
  coveragePathIgnorePatterns: [
    "/node_modules/",
    ".test.ts",
    ".stories.ts",
    ".d.ts",
  ],
  testEnvironment: "./custom-jsdom",
  clearMocks: true,
  setupFiles: ["./jest.setup.js"],
  setupFilesAfterEnv: ["./jest.after-env.ts"],
  coverageReporters: ["lcovonly", "text-summary", "html"],
  maxWorkers: 4,
  globals: {
    "ts-jest": {
      isolatedModules: true,
    },
  },
  testTimeout: 10000,
};
