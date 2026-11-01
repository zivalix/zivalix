defmodule Zivalix.MixProject do
  use Mix.Project

  def project do
    [
      app: :zivalix,
      version: "0.1.0",
      description: "Official reserved package for ZIVALIX Agency. Defensive registration.",
      package: package(),
      deps: [],
      build_embedded: Mix.env() == :prod,
      start_permanent: Mix.env() == :prod
    ]
  end

  def application do
    [extra_applications: []]
  end

  defp package do
    [
      licenses: ["ISC"],
      links: %{"Homepage" => "https://zivalix.com", "GitHub" => "https://github.com/zivalix/zivalix"}
    ]
  end
end
