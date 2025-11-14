## Development Process

### Initial Research & Setup
- **30-40 minutes**: I made some research on Go and Svelte fundamentals, understanding the frameworks and their syntax
- Setup of development environment

### My Implementation Process

**First Step: Core Requirements (40-50 minutes)**
- Implemented basic backend functionality:
  - GET endpoint to list todos
  - POST endpoint to add todos
  - In-memory storage system
- Completed frontend form submission:
  - Connected form to backend API
  - Implemented basic error handling
  - Form validation

**Second Step: Enhancements & Improvements (70-80 minutes)**
- I did some additional research on best practices for Go and Svelte to have a better understanding and wanted to make improvements in order to practice.
- Implemented additional features:
  - Priority system (urgent, medium, low) with sorting
  - Todo completion status with checkboxes
  - Edit and delete functionality
  - Dark/light mode theme toggle
  - Improved UI/UX
- Code quality improvements:
  - Error handling and validation
  - Code organization and helper functions
  - Accessibility improvements (ARIA labels) (This is just something I like to add!)
  - Code comments and documentation

**Third Step: Testing**
- Before submitting my code, I tested:
  - Backend API endpoints
  - Frontend form validations
  - Priority system (selection, display, sorting)
  - Completion status
  - Edit functionality
  - Delete functionality
  - Theme toggle

## What I Would Do If I Had More Time
- After making a bit more research and also using AI as a tool to review and point possible weak sides of the code, these are some of the things I took a note of:
### Backend Improvements
- **Thread Safety**: Implement `sync.Mutex` to protect shared state (`todos` slice and `nextID`) from concurrent access, preventing race conditions in a production environment
- **Constants**: Replace magic strings (priority values) with constants for better maintainability and type safety
- **Logging**: Add structured logging for requests, errors, and important events to aid debugging and monitoring
- **Configuration**: Make CORS origin configurable via environment variables instead of hardcoded `*` for production security
- **Error Handling**: Implement more granular error types and standardized error response format
- **Testing**: Add unit tests for handlers and validation functions, plus integration tests for API endpoints

### Frontend Improvements
- **Environment Configuration**: Move API URL to environment variables or config file for different environments
- **User Experience**: Replace `alert()` with toast notifications for better UX and non-blocking error messages
- **Loading States**: Add loading indicators and disable buttons during API calls to prevent duplicate submissions
- **State Management**: Improve state synchronization in Todo component to handle external prop changes better
- **Error Recovery**: Implement retry logic for failed API calls with exponential backoff
- **Testing**: Add component tests and E2E tests for critical user flows
- **Desing**: Improve UI/UX and make a better design.

### Code Quality
- **Documentation**: Add more comprehensive inline documentation and API documentation
- **Type Safety**: Further strengthen TypeScript types and add runtime validation where needed
