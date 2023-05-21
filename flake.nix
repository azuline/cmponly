{
  description = "cmponly development environment";

  inputs = {
    nixpkgs.url = github:nixos/nixpkgs/nixos-unstable;
    flake-utils.url = github:numtide/flake-utils;
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system}; in rec {
        devShell = pkgs.mkShell {
          buildInputs = [
            (pkgs.buildEnv {
              name = "cmponly-env";
              paths = with pkgs; [
                gnumake
                golangci-lint
                go_1_18
                gopls
                gopkgs
                go-outline
                delve
                gotools
              ];
            })
          ];
          shellHook = ''
            # Isolate build stuff to this repo's directory.
            export CMPONLY_ROOT="$(pwd)"
            export CMPONLY_CACHE_ROOT="$(pwd)/.cache"

            export GOCACHE="$CMPONLY_CACHE_ROOT/go/cache"
            export GOENV="$CMPONLY_CACHE_ROOT/go/env"
            export GOPATH="$CMPONLY_CACHE_ROOT/go/path"
            export GOMODCACHE="$GOPATH/pkg/mod"
            export GOROOT=
            export PATH=$(go env GOPATH)/bin:$PATH
          '';
        };
      }
    );
}
