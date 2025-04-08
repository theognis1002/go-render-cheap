# Render Services Manager

A command-line tool to manage Render.com services. This tool allows you to suspend and resume multiple Render services in bulk using the Render API.

## Features

- Suspend and resume Render services
- Bulk operations on multiple services at once
- Detailed error reporting

## Prerequisites

- Go 1.x or higher
- A Render API key
- Service IDs from your Render dashboard

## Installation

```bash
git clone https://github.com/yourusername/go-render-services.git
cd go-render-services
go build -o render-manager
```

This will create a binary named `render-manager` in your current directory. You can move it to a location in your PATH if desired:

```bash
# Optional: Move to a directory in your PATH (Linux/macOS)
sudo mv render-manager /usr/local/bin/
```

## Usage

1. Set your Render API key as an environment variable:

```bash
export RENDER_API_KEY="your_render_api_key"
```

2. Set your service IDs as a comma-separated list:

```bash
export RENDER_SERVICE_IDS="srv-123,srv-456,srv-789"
```

3. Run using either the binary or `go run`:

Using the binary:

```bash
# To suspend services
./render-manager suspend

# To resume services
./render-manager unsuspend
```

Or using `go run`:

```bash
# To suspend services
go run main.go suspend

# To resume services
go run main.go unsuspend
```

### Example

```bash
# Suspend multiple services
export RENDER_SERVICE_IDS="srv-abc123,srv-def456"
./render-manager suspend

# Resume the same services
./render-manager unsuspend
```

## Error Handling

The script will:

- Continue processing remaining services if one fails
- Display detailed error messages for each failed operation
- Show success messages for each successful operation

## Environment Variables

| Variable             | Description                         | Required |
| -------------------- | ----------------------------------- | -------- |
| `RENDER_API_KEY`     | Your Render API key                 | Yes      |
| `RENDER_SERVICE_IDS` | Comma-separated list of service IDs | Yes      |

## Notes

- All operations are performed sequentially
- Each request has a 10-second timeout

## License

MIT License (or specify your chosen license)
