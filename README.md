# ddd-ct - Domain-Driven Design Collaboration Tool (DDD-CT)

### Quickstart
* Windows
    * `docker-compose up --build`
    * `go run ./cmd/webserver`

### Overview
This tool aims to create a more structured approach to the genesis of domain-driven design (DDD) for a given project.

### Core Features:

* Idea Boards: Similar to digital whiteboards, these are spaces where participants can post ideas. However, they're organized into DDD-specific categories such as "Domains," "Subdomains," "Bounded Contexts," and "Ubiquitous Language."
* Context Mapping Workshop: An interactive module that guides the team through creating context maps. This helps in identifying bounded contexts and the relationships between them, which is a key component of DDD.
* Ubiquitous Language Glossary: A collaborative glossary where terms are defined and refined. This ensures that all team members share a common language, reducing confusion and aligning understanding.
* Modeling Workshops: Facilitated sessions within the app that guide the team through exercises in Event Storming or Example Mapping to identify domain events, commands, aggregates, and entities.
* Priority Voting: To keep discussions focused, the tool includes a feature for voting on ideas or models that should be prioritized. This helps in focusing the discussion on what matters most to the project.

#### Bounded Contexts:
* User Management: Handles user registration, authentication, and profile management. It focuses on user identity and access control.
* Project Collaboration: Encompasses the collaboration and interaction among users within projects. This includes project creation, management, and user roles within those projects.
* Domain Modeling: Central to the application, this BC deals with the creation and management of domain models, context maps, and ubiquitous language glossaries specific to each project.
* Integration: Manages the integration with external systems and tools, such as project management software or version control systems.

#### Domains and Subdomains:
* Core Domain: Domain Modeling BC, as it's the unique proposition of the tool.
* Supporting Subdomains: User Management and Project Collaboration, which support the core functionality but are not the unique selling points.
* Generic Subdomains: Integration with other tools and services, which can be solved with off-the-shelf solutions.

#### Entities:
* Project: Represents a collaboration project with its unique identifier, name, and associated domain models.
* Domain Model: Represents a specific domain model within a project. It could be an aggregate root for entities, value objects, events, and bounded contexts within the model.

#### Value Objects:
* Ubiquitous Language Term: Represents terms defined in the project's ubiquitous language glossary.
* Domain Event: Represents a significant event within the domain modeling process.

#### Aggregates:
* Domain Model Aggregate: Encompasses the Domain Model entity as the root, along with nested entities like Entities, Value Objects, and Domain Events as part of the model.

#### Domain Services:
* Modeling Service: Offers operations for creating and modifying domain models, including adding or removing elements from a model.
* Collaboration Service: Facilitates real-time collaboration among users, such as live editing of models and discussions.

#### Repositories:
* Project Repository: Manages persistence and retrieval of Project aggregates.
* Domain Model Repository: Manages persistence and retrieval of Domain Model aggregates.

#### Application Services:
* Project Management Service: Handles creation, deletion, and updates to projects.
* Access Control Service: Manages user permissions within a project.

#### Infrastructure Services:
* Authentication Service: Validates user credentials and manages sessions.
* Integration Adapters: Connects to external project management tools or version control systems.
