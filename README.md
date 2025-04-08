# Render Services Manager

A command-line tool to manage Render.com services and databases. This tool allows you to suspend and resume/restart multiple Render services and databases in bulk using the Render API.

## Features

- Suspend and resume Render services
- Suspend and resume Render databases
- Bulk operations on multiple services/databases at once
- Automatic detection of database vs service IDs
- Detailed error reporting

## Prerequisites

- Go 1.x or higher
- A Render API key
- Service/Database IDs from your Render dashboard

## Installation

```bash
git clone https://github.com/yourusername/go-render-services.git
cd go-render-services
go build
```

## Usage

1. Set your Render API key as an environment variable:

```bash
export RENDER_API_KEY="your_render_api_key"
```

2. Set your service/database IDs as a comma-separated list:

```bash
export RENDER_SERVICE_IDS="srv-123,dpg-456,srv-789"
```

3. Run the script with either `suspend` or `unsuspend`:

```bash
# To suspend services/databases
go run main.go suspend

# To resume services/restart databases
go run main.go unsuspend
```

### Example

```bash
# Suspend multiple services and databases
export RENDER_SERVICE_IDS="srv-abc123,dpg-xyz789,srv-def456"
go run main.go suspend

# Resume the same services and databases
go run main.go unsuspend
```

## Service ID Types

The script automatically detects the type of resource based on the ID prefix:

- Database IDs start with `dpg-`
- Service IDs typically start with `srv-`

## Error Handling

The script will:

- Continue processing remaining services if one fails
- Display detailed error messages for each failed operation
- Show success messages for each successful operation

## Environment Variables

| Variable             | Description                                  | Required |
| -------------------- | -------------------------------------------- | -------- |
| `RENDER_API_KEY`     | Your Render API key                          | Yes      |
| `RENDER_SERVICE_IDS` | Comma-separated list of service/database IDs | Yes      |

## Notes

- The script uses different API endpoints for databases and services
- For services, "unsuspend" actually performs a restart operation
- For databases, "unsuspend" performs a resume operation
- All operations are performed sequentially
- Each request has a 10-second timeout

## License

MIT License (or specify your chosen license)
