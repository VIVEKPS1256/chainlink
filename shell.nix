{pkgs}:
with pkgs; let
  go = go_1_21;
  postgresql = postgresql_14;
  nodejs = nodejs-18_x;
  nodePackages = pkgs.nodePackages.override {inherit nodejs;};
  pnpm = pnpm_9;

  mkShell' = mkShell.override {
    # The current nix default sdk for macOS fails to compile go projects, so we use a newer one for now.
    stdenv =
      if stdenv.isDarwin
      then overrideSDK stdenv "11.0"
      else stdenv;
  };
in
  mkShell' {
    nativeBuildInputs =
      [
        go
        nur.repos.goreleaser.goreleaser-pro
        postgresql

        python3
        python3Packages.pip
        protobuf
        protoc-gen-go
        protoc-gen-go-grpc

        foundry-bin

        curl
        nodejs
        pnpm
        # TODO: compiler / gcc for secp compilation
        go-ethereum # geth
        go-mockery

        # tooling
        gotools
        gopls
        delve
        golangci-lint
        github-cli
        jq

        # gofuzz
      ]
      ++ lib.optionals stdenv.isLinux [
        # some dependencies needed for node-gyp on pnpm install
        pkg-config
        libudev-zero
        libusb1
      ];

    # We expect the user to install the cross compile toolchain via
    #
    # brew tap messense/macos-cross-toolchains 
    # brew install aarch64-unknown-linux-gnu
    #
    # TODO: Ideally we manage this in nix too, but I could not figure out how to
    # do so
    shellHook = ''
        ${if stdenv.isDarwin then "
          export CC=/opt/homebrew/Cellar/aarch64-unknown-linux-gnu/13.3.0/bin/aarch64-linux-gnu-gcc
          export CXX=/opt/homebrew/Cellar/aarch64-unknown-linux-gnu/13.3.0/bin/aarch64-linux-gnu-g++
        " else ""}
      '';

    GOROOT = "${go}/share/go";
    PGDATA = "db";
    CL_DATABASE_URL = "postgresql://chainlink:chainlink@localhost:5432/chainlink_test?sslmode=disable";
  }
