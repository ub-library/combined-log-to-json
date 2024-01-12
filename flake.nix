{
  description = "A basic flake with a shell";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-23.11-darwin";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        combined-log-to-json = pkgs.callPackage ./package.nix {};
      in
      {
        inherit combined-log-to-json;

        defaultPackage = combined-log-to-json;

        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.bashInteractive
            pkgs.go
          ];
        };
      });
}
