{
  description = "A basic flake with a shell";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-23.11-darwin";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        defaultPackage = pkgs.buildGoModule rec {
          name = "combined-log-to-json";
          src = builtins.path {
            path = ./.;
            name = "${name}-src";
          };
          vendorHash = null;
        };
        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.bashInteractive
            pkgs.go
          ];
        };
      });
}
