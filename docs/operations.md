
# actions


## actions add-selected-repo-to-org-secret

https://developer.github.com/v3/actions/secrets/#add-selected-repository-to-an-organization-secret

Adds a repository to an organization secret when the `visibility` for repository access is set to `selected`. The visibility is set when you [Create or update an organization secret](https://developer.github.com/v3/actions/secrets/#create-or-update-an-organization-secret). You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| repository_id | __Required__ repository_id parameter |
| secret_name | __Required__ secret_name parameter |

## actions cancel-workflow-run

https://developer.github.com/v3/actions/workflow-runs/#cancel-a-workflow-run

Cancels a workflow run using its `id`. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `actions:write` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |

## actions create-or-update-org-secret

https://developer.github.com/v3/actions/secrets/#create-or-update-an-organization-secret

Creates or updates an organization secret with an encrypted value. Encrypt your secret using
[LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages). You must authenticate using an access
token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to
use this endpoint.

#### Example encrypting a secret using Node.js

Encrypt your secret using the [tweetsodium](https://github.com/github/tweetsodium) library.

```
const sodium = require('tweetsodium');

const key = "base64-encoded-public-key";
const value = "plain-text-secret";

// Convert the message and key to Uint8Array's (Buffer implements that interface)
const messageBytes = Buffer.from(value);
const keyBytes = Buffer.from(key, 'base64');

// Encrypt using LibSodium.
const encryptedBytes = sodium.seal(messageBytes, keyBytes);

// Base64 the encrypted secret
const encrypted = Buffer.from(encryptedBytes).toString('base64');

console.log(encrypted);
```


#### Example encrypting a secret using Python

Encrypt your secret using [pynacl](https://pynacl.readthedocs.io/en/stable/public/#nacl-public-sealedbox) with Python 3.

```
from base64 import b64encode
from nacl import encoding, public

def encrypt(public_key: str, secret_value: str) -> str:
  """Encrypt a Unicode string using the public key."""
  public_key = public.PublicKey(public_key.encode("utf-8"), encoding.Base64Encoder())
  sealed_box = public.SealedBox(public_key)
  encrypted = sealed_box.encrypt(secret_value.encode("utf-8"))
  return b64encode(encrypted).decode("utf-8")
```

#### Example encrypting a secret using C#

Encrypt your secret using the [Sodium.Core](https://www.nuget.org/packages/Sodium.Core/) package.

```
var secretValue = System.Text.Encoding.UTF8.GetBytes("mySecret");
var publicKey = Convert.FromBase64String("2Sg8iYjAxxmI2LvUXpJjkYrMxURPc8r+dB7TJyvvcCU=");

var sealedPublicKeyBox = Sodium.SealedPublicKeyBox.Create(secretValue, publicKey);

Console.WriteLine(Convert.ToBase64String(sealedPublicKeyBox));
```

#### Example encrypting a secret using Ruby

Encrypt your secret using the [rbnacl](https://github.com/RubyCrypto/rbnacl) gem.

```ruby
require "rbnacl"
require "base64"

key = Base64.decode64("+ZYvJDZMHUfBkJdyq5Zm9SKqeuBQ4sj+6sfjlH4CgG0=")
public_key = RbNaCl::PublicKey.new(key)

box = RbNaCl::Boxes::Sealed.from_public_key(public_key)
encrypted_secret = box.encrypt("my_secret")

# Print the base64 encoded secret
puts Base64.strict_encode64(encrypted_secret)
```

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| secret_name | __Required__ secret_name parameter |
| encrypted_value | Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get an organization public key](https://developer.github.com/v3/actions/secrets/#get-an-organization-public-key) endpoint. |
| key_id | ID of the key you used to encrypt the secret. |
| selected_repository_ids | An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the `visibility` is set to `selected`. You can manage the list of selected repositories using the [List selected repositories for an organization secret](https://developer.github.com/v3/actions/secrets/#list-selected-repositories-for-an-organization-secret), [Set selected repositories for an organization secret](https://developer.github.com/v3/actions/secrets/#set-selected-repositories-for-an-organization-secret), and [Remove selected repository from an organization secret](https://developer.github.com/v3/actions/secrets/#remove-selected-repository-from-an-organization-secret) endpoints. |
| visibility | Configures the access that repositories have to the organization secret. Can be one of:  <br>\- `all` - All repositories in an organization can access the secret.  <br>\- `private` - Private repositories in an organization can access the secret.  <br>\- `selected` - Only specific repositories can access the secret. |

## actions create-or-update-repo-secret

https://developer.github.com/v3/actions/secrets/#create-or-update-a-repository-secret

Creates or updates a repository secret with an encrypted value. Encrypt your secret using
[LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages). You must authenticate using an access
token with the `repo` scope to use this endpoint. GitHub Apps must have the `secrets` repository permission to use
this endpoint.

#### Example encrypting a secret using Node.js

Encrypt your secret using the [tweetsodium](https://github.com/github/tweetsodium) library.

```
const sodium = require('tweetsodium');

const key = "base64-encoded-public-key";
const value = "plain-text-secret";

// Convert the message and key to Uint8Array's (Buffer implements that interface)
const messageBytes = Buffer.from(value);
const keyBytes = Buffer.from(key, 'base64');

// Encrypt using LibSodium.
const encryptedBytes = sodium.seal(messageBytes, keyBytes);

// Base64 the encrypted secret
const encrypted = Buffer.from(encryptedBytes).toString('base64');

console.log(encrypted);
```


#### Example encrypting a secret using Python

Encrypt your secret using [pynacl](https://pynacl.readthedocs.io/en/stable/public/#nacl-public-sealedbox) with Python 3.

```
from base64 import b64encode
from nacl import encoding, public

def encrypt(public_key: str, secret_value: str) -> str:
  """Encrypt a Unicode string using the public key."""
  public_key = public.PublicKey(public_key.encode("utf-8"), encoding.Base64Encoder())
  sealed_box = public.SealedBox(public_key)
  encrypted = sealed_box.encrypt(secret_value.encode("utf-8"))
  return b64encode(encrypted).decode("utf-8")
```

#### Example encrypting a secret using C#

Encrypt your secret using the [Sodium.Core](https://www.nuget.org/packages/Sodium.Core/) package.

```
var secretValue = System.Text.Encoding.UTF8.GetBytes("mySecret");
var publicKey = Convert.FromBase64String("2Sg8iYjAxxmI2LvUXpJjkYrMxURPc8r+dB7TJyvvcCU=");

var sealedPublicKeyBox = Sodium.SealedPublicKeyBox.Create(secretValue, publicKey);

Console.WriteLine(Convert.ToBase64String(sealedPublicKeyBox));
```

#### Example encrypting a secret using Ruby

Encrypt your secret using the [rbnacl](https://github.com/RubyCrypto/rbnacl) gem.

```ruby
require "rbnacl"
require "base64"

key = Base64.decode64("+ZYvJDZMHUfBkJdyq5Zm9SKqeuBQ4sj+6sfjlH4CgG0=")
public_key = RbNaCl::PublicKey.new(key)

box = RbNaCl::Boxes::Sealed.from_public_key(public_key)
encrypted_secret = box.encrypt("my_secret")

# Print the base64 encoded secret
puts Base64.strict_encode64(encrypted_secret)
```

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| secret_name | __Required__ secret_name parameter |
| encrypted_value | Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get a repository public key](https://developer.github.com/v3/actions/secrets/#get-a-repository-public-key) endpoint. |
| key_id | ID of the key you used to encrypt the secret. |

## actions create-registration-token-for-org

https://developer.github.com/v3/actions/self-hosted-runners/#create-a-registration-token-for-an-organization

**Warning:** The self-hosted runners API for organizations is currently in public beta and subject to change.


Returns a token that you can pass to the `config` script. The token expires after one hour. You must authenticate
using an access token with the `admin:org` scope to use this endpoint.

#### Example using registration token

Configure your self-hosted runner, replacing `TOKEN` with the registration token provided by this endpoint.

```
./config.sh --url https://github.com/octo-org --token TOKEN
```

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## actions create-registration-token-for-repo

https://developer.github.com/v3/actions/self-hosted-runners/#create-a-registration-token-for-a-repository

Returns a token that you can pass to the `config` script. The token expires after one hour. You must authenticate
using an access token with the `repo` scope to use this endpoint.

#### Example using registration token
 
Configure your self-hosted runner, replacing `TOKEN` with the registration token provided by this endpoint.

```
./config.sh --url https://github.com/octo-org/octo-repo-artifacts --token TOKEN
```

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## actions create-remove-token-for-org

https://developer.github.com/v3/actions/self-hosted-runners/#create-a-remove-token-for-an-organization

**Warning:** The self-hosted runners API for organizations is currently in public beta and subject to change.


Returns a token that you can pass to the `config` script to remove a self-hosted runner from an organization. The
token expires after one hour. You must authenticate using an access token with the `admin:org` scope to use this
endpoint.

#### Example using remove token

To remove your self-hosted runner from an organization, replace `TOKEN` with the remove token provided by this
endpoint.

```
./config.sh remove --token TOKEN
```

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## actions create-remove-token-for-repo

https://developer.github.com/v3/actions/self-hosted-runners/#create-a-remove-token-for-a-repository

Returns a token that you can pass to remove a self-hosted runner from a repository. The token expires after one hour.
You must authenticate using an access token with the `repo` scope to use this endpoint.

#### Example using remove token
 
To remove your self-hosted runner from a repository, replace TOKEN with the remove token provided by this endpoint.

```
./config.sh remove --token TOKEN
```

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## actions create-workflow-dispatch

https://developer.github.com/v3/actions/workflows/#create-a-workflow-dispatch-event

You can use this endpoint to manually trigger a GitHub Actions workflow run. You can also replace `{workflow_id}` with the workflow file name. For example, you could use `main.yml`.

You must configure your GitHub Actions workflow to run when the [`workflow_dispatch` webhook](/developers/webhooks-and-events/webhook-events-and-payloads#workflow_dispatch) event occurs. The `inputs` are configured in the workflow file. For more information about how to configure the `workflow_dispatch` event in the workflow file, see "[Events that trigger workflows](/actions/reference/events-that-trigger-workflows#workflow_dispatch)."

You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `actions:write` permission to use this endpoint. For more information, see "[Creating a personal access token for the command line](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line)."

### parameters


| name | description |
|------|-------------|
| ref | __Required__ The reference of the workflow run. The reference can be a branch, tag, or a commit SHA. |
| repo | __Required__ repository in OWNER/REPO form |
| workflow_id | __Required__  |
| inputs | Input keys and values configured in the workflow file. The maximum number of properties is 10. Any default properties configured in the workflow file will be used when `inputs` are omitted. |

## actions delete-artifact

https://developer.github.com/v3/actions/artifacts/#delete-an-artifact

Deletes an artifact for a workflow run. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `actions:write` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| artifact_id | __Required__ artifact_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## actions delete-org-secret

https://developer.github.com/v3/actions/secrets/#delete-an-organization-secret

Deletes a secret in an organization using the secret name. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| secret_name | __Required__ secret_name parameter |

## actions delete-repo-secret

https://developer.github.com/v3/actions/secrets/#delete-a-repository-secret

Deletes a secret in a repository using the secret name. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `secrets` repository permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| secret_name | __Required__ secret_name parameter |

## actions delete-self-hosted-runner-from-org

https://developer.github.com/v3/actions/self-hosted-runners/#delete-a-self-hosted-runner-from-an-organization

**Warning:** The self-hosted runners API for organizations is currently in public beta and subject to change.

Forces the removal of a self-hosted runner from an organization. You can use this endpoint to completely remove the runner when the machine you were using no longer exists. You must authenticate using an access token with the `admin:org` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| runner_id | __Required__ runner_id parameter |

## actions delete-self-hosted-runner-from-repo

https://developer.github.com/v3/actions/self-hosted-runners/#delete-a-self-hosted-runner-from-a-repository

Forces the removal of a self-hosted runner from a repository. You can use this endpoint to completely remove the runner when the machine you were using no longer exists. You must authenticate using an access token with the `repo` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| runner_id | __Required__ runner_id parameter |

## actions delete-workflow-run

https://developer.github.com/v3/actions/workflow-runs/#delete-a-workflow-run

Delete a specific workflow run. Anyone with write access to the repository can use this endpoint. If the repository is
private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:write` permission to use
this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |

## actions delete-workflow-run-logs

https://developer.github.com/v3/actions/workflow-runs/#delete-workflow-run-logs

Deletes all logs for a workflow run. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `actions:write` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |

## actions download-artifact

https://developer.github.com/v3/actions/artifacts/#download-an-artifact

Gets a redirect URL to download an archive for a repository. This URL expires after 1 minute. Look for `Location:` in
the response header to find the URL for the download. The `:archive_format` must be `zip`. Anyone with read access to
the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.
GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| archive_format | __Required__ archive_format parameter |
| artifact_id | __Required__ artifact_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## actions download-job-logs-for-workflow-run

https://developer.github.com/v3/actions/workflow-jobs/#download-job-logs-for-a-workflow-run

Gets a redirect URL to download a plain text file of logs for a workflow job. This link expires after 1 minute. Look
for `Location:` in the response header to find the URL for the download. Anyone with read access to the repository can
use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must
have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| job_id | __Required__ job_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## actions download-workflow-run-logs

https://developer.github.com/v3/actions/workflow-runs/#download-workflow-run-logs

Gets a redirect URL to download an archive of log files for a workflow run. This link expires after 1 minute. Look for
`Location:` in the response header to find the URL for the download. Anyone with read access to the repository can use
this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have
the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |

## actions get-artifact

https://developer.github.com/v3/actions/artifacts/#get-an-artifact

Gets a specific artifact for a workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| artifact_id | __Required__ artifact_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## actions get-job-for-workflow-run

https://developer.github.com/v3/actions/workflow-jobs/#get-a-job-for-a-workflow-run

Gets a specific job in a workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| job_id | __Required__ job_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## actions get-org-public-key

https://developer.github.com/v3/actions/secrets/#get-an-organization-public-key

Gets your public key, which you need to encrypt secrets. You need to encrypt a secret before you can create or update secrets. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## actions get-org-secret

https://developer.github.com/v3/actions/secrets/#get-an-organization-secret

Gets a single organization secret without revealing its encrypted value. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| secret_name | __Required__ secret_name parameter |

## actions get-repo-public-key

https://developer.github.com/v3/actions/secrets/#get-a-repository-public-key

Gets your public key, which you need to encrypt secrets. You need to encrypt a secret before you can create or update secrets. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `secrets` repository permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## actions get-repo-secret

https://developer.github.com/v3/actions/secrets/#get-a-repository-secret

Gets a single repository secret without revealing its encrypted value. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `secrets` repository permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| secret_name | __Required__ secret_name parameter |

## actions get-self-hosted-runner-for-org

https://developer.github.com/v3/actions/self-hosted-runners/#get-a-self-hosted-runner-for-an-organization

**Warning:** The self-hosted runners API for organizations is currently in public beta and subject to change.

Gets a specific self-hosted runner for an organization. You must authenticate using an access token with the `admin:org` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| runner_id | __Required__ runner_id parameter |

## actions get-self-hosted-runner-for-repo

https://developer.github.com/v3/actions/self-hosted-runners/#get-a-self-hosted-runner-for-a-repository

Gets a specific self-hosted runner. You must authenticate using an access token with the `repo` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| runner_id | __Required__ runner_id parameter |

## actions get-workflow

https://developer.github.com/v3/actions/workflows/#get-a-workflow

Gets a specific workflow. You can also replace `:workflow_id` with `:workflow_file_name`. For example, you could use `main.yml`. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| workflow_id | __Required__  |

## actions get-workflow-run

https://developer.github.com/v3/actions/workflow-runs/#get-a-workflow-run

Gets a specific workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |

## actions get-workflow-run-usage

https://developer.github.com/v3/actions/workflow-runs/#get-workflow-run-usage

**Warning:** This GitHub Actions usage endpoint is currently in public beta and subject to change. For more information, see "[GitHub Actions API workflow usage](https://developer.github.com/changes/2020-05-15-actions-api-workflow-usage)."

Gets the number of billable minutes and total run time for a specific workflow run. Billable minutes only apply to workflows in private repositories that use GitHub-hosted runners. Usage is listed for each GitHub-hosted runner operating system in milliseconds. Any job re-runs are also included in the usage. The usage does not include the multiplier for macOS and Windows runners and is not rounded up to the nearest whole minute. For more information, see "[Managing billing for GitHub Actions](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-actions)".

Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |

## actions get-workflow-usage

https://developer.github.com/v3/actions/workflows/#get-workflow-usage

**Warning:** This GitHub Actions usage endpoint is currently in public beta and subject to change. For more information, see "[GitHub Actions API workflow usage](https://developer.github.com/changes/2020-05-15-actions-api-workflow-usage)."

Gets the number of billable minutes used by a specific workflow during the current billing cycle. Billable minutes only apply to workflows in private repositories that use GitHub-hosted runners. Usage is listed for each GitHub-hosted runner operating system in milliseconds. Any job re-runs are also included in the usage. The usage does not include the multiplier for macOS and Windows runners and is not rounded up to the nearest whole minute. For more information, see "[Managing billing for GitHub Actions](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-actions)".

You can also replace `:workflow_id` with `:workflow_file_name`. For example, you could use `main.yml`. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| workflow_id | __Required__  |

## actions list-artifacts-for-repo

https://developer.github.com/v3/actions/artifacts/#list-artifacts-for-a-repository

Lists all artifacts for a repository. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## actions list-jobs-for-workflow-run

https://developer.github.com/v3/actions/workflow-jobs/#list-jobs-for-a-workflow-run

Lists jobs for a workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://developer.github.com/v3/#parameters).

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |
| filter | Filters jobs by their `completed_at` timestamp. Can be one of:  <br>\* `latest`: Returns jobs from the most recent execution of the workflow run.  <br>\* `all`: Returns all jobs for a workflow run, including from old executions of the workflow run. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## actions list-org-secrets

https://developer.github.com/v3/actions/secrets/#list-organization-secrets

Lists all secrets available in an organization without revealing their encrypted values. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## actions list-repo-secrets

https://developer.github.com/v3/actions/secrets/#list-repository-secrets

Lists all secrets available in a repository without revealing their encrypted values. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `secrets` repository permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## actions list-repo-workflows

https://developer.github.com/v3/actions/workflows/#list-repository-workflows

Lists the workflows in a repository. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## actions list-runner-applications-for-org

https://developer.github.com/v3/actions/self-hosted-runners/#list-runner-applications-for-an-organization

**Warning:** The self-hosted runners API for organizations is currently in public beta and subject to change.

Lists binaries for the runner application that you can download and run. You must authenticate using an access token with the `admin:org` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## actions list-runner-applications-for-repo

https://developer.github.com/v3/actions/self-hosted-runners/#list-runner-applications-for-a-repository

Lists binaries for the runner application that you can download and run. You must authenticate using an access token with the `repo` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## actions list-selected-repos-for-org-secret

https://developer.github.com/v3/actions/secrets/#list-selected-repositories-for-an-organization-secret

Lists all repositories that have been selected when the `visibility` for repository access to a secret is set to `selected`. You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| secret_name | __Required__ secret_name parameter |

## actions list-self-hosted-runners-for-org

https://developer.github.com/v3/actions/self-hosted-runners/#list-self-hosted-runners-for-an-organization

**Warning:** The self-hosted runners API for organizations is currently in public beta and subject to change.

Lists all self-hosted runners for an organization. You must authenticate using an access token with the `admin:org` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## actions list-self-hosted-runners-for-repo

https://developer.github.com/v3/actions/self-hosted-runners/#list-self-hosted-runners-for-a-repository

Lists all self-hosted runners for a repository. You must authenticate using an access token with the `repo` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## actions list-workflow-run-artifacts

https://developer.github.com/v3/actions/artifacts/#list-workflow-run-artifacts

Lists artifacts for a workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## actions list-workflow-runs

https://developer.github.com/v3/actions/workflow-runs/#list-workflow-runs

List all workflow runs for a workflow. You can also replace `:workflow_id` with `:workflow_file_name`. For example, you could use `main.yml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://developer.github.com/v3/#parameters).

Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| workflow_id | __Required__  |
| actor | Returns someone's workflow runs. Use the login for the user who created the `push` associated with the check suite or workflow run. |
| branch | Returns workflow runs associated with a branch. Use the name of the branch of the `push`. |
| event | Returns workflow run triggered by the event you specify. For example, `push`, `pull_request` or `issue`. For more information, see "[Events that trigger workflows](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows)." |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| status | Returns workflow runs associated with the check run `status` or `conclusion` you specify. For example, a conclusion can be `success` or a status can be `completed`. For more information, see the `status` and `conclusion` options available in "[Create a check run](https://developer.github.com/v3/checks/runs/#create-a-check-run)." |

## actions list-workflow-runs-for-repo

https://developer.github.com/v3/actions/workflow-runs/#list-workflow-runs-for-a-repository

Lists all workflow runs for a repository. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://developer.github.com/v3/#parameters).

Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| actor | Returns someone's workflow runs. Use the login for the user who created the `push` associated with the check suite or workflow run. |
| branch | Returns workflow runs associated with a branch. Use the name of the branch of the `push`. |
| event | Returns workflow run triggered by the event you specify. For example, `push`, `pull_request` or `issue`. For more information, see "[Events that trigger workflows](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows)." |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| status | Returns workflow runs associated with the check run `status` or `conclusion` you specify. For example, a conclusion can be `success` or a status can be `completed`. For more information, see the `status` and `conclusion` options available in "[Create a check run](https://developer.github.com/v3/checks/runs/#create-a-check-run)." |

## actions re-run-workflow

https://developer.github.com/v3/actions/workflow-runs/#re-run-a-workflow

Re-runs your workflow run using its `id`. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `actions:write` permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| run_id | __Required__  |

## actions remove-selected-repo-from-org-secret

https://developer.github.com/v3/actions/secrets/#remove-selected-repository-from-an-organization-secret

Removes a repository from an organization secret when the `visibility` for repository access is set to `selected`. The visibility is set when you [Create or update an organization secret](https://developer.github.com/v3/actions/secrets/#create-or-update-an-organization-secret). You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| repository_id | __Required__ repository_id parameter |
| secret_name | __Required__ secret_name parameter |

## actions set-selected-repos-for-org-secret

https://developer.github.com/v3/actions/secrets/#set-selected-repositories-for-an-organization-secret

Replaces all repositories for an organization secret when the `visibility` for repository access is set to `selected`. The visibility is set when you [Create or update an organization secret](https://developer.github.com/v3/actions/secrets/#create-or-update-an-organization-secret). You must authenticate using an access token with the `admin:org` scope to use this endpoint. GitHub Apps must have the `secrets` organization permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| secret_name | __Required__ secret_name parameter |
| selected_repository_ids | An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the `visibility` is set to `selected`. You can add and remove individual repositories using the [Set selected repositories for an organization secret](https://developer.github.com/v3/actions/secrets/#set-selected-repositories-for-an-organization-secret) and [Remove selected repository from an organization secret](https://developer.github.com/v3/actions/secrets/#remove-selected-repository-from-an-organization-secret) endpoints. |

# activity


## activity check-repo-is-starred-by-authenticated-user

https://developer.github.com/v3/activity/starring/#check-if-a-repository-is-starred-by-the-authenticated-user



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## activity delete-repo-subscription

https://developer.github.com/v3/activity/watching/#delete-a-repository-subscription

This endpoint should only be used to stop watching a repository. To control whether or not you wish to receive notifications from a repository, [set the repository's subscription manually](https://developer.github.com/v3/activity/watching/#set-a-repository-subscription).

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## activity delete-thread-subscription

https://developer.github.com/v3/activity/notifications/#delete-a-thread-subscription

Mutes all future notifications for a conversation until you comment on the thread or get an **@mention**. If you are watching the repository of the thread, you will still receive notifications. To ignore future notifications for a repository you are watching, use the [Set a thread subscription](https://developer.github.com/v3/activity/notifications/#set-a-thread-subscription) endpoint and set `ignore` to `true`.

### parameters


| name | description |
|------|-------------|
| thread_id | __Required__ thread_id parameter |

## activity get-feeds

https://developer.github.com/v3/activity/feeds/#get-feeds

GitHub provides several timeline resources in [Atom](http://en.wikipedia.org/wiki/Atom_(standard)) format. The Feeds API lists all the feeds available to the authenticated user:

*   **Timeline**: The GitHub global public timeline
*   **User**: The public timeline for any user, using [URI template](https://developer.github.com/v3/#hypermedia)
*   **Current user public**: The public timeline for the authenticated user
*   **Current user**: The private timeline for the authenticated user
*   **Current user actor**: The private timeline for activity created by the authenticated user
*   **Current user organizations**: The private timeline for the organizations the authenticated user is a member of.
*   **Security advisories**: A collection of public announcements that provide information about security-related vulnerabilities in software on GitHub.

**Note**: Private feeds are only returned when [authenticating via Basic Auth](https://developer.github.com/v3/#basic-authentication) since current feed URIs use the older, non revocable auth tokens.

## activity get-repo-subscription

https://developer.github.com/v3/activity/watching/#get-a-repository-subscription



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## activity get-thread

https://developer.github.com/v3/activity/notifications/#get-a-thread



### parameters


| name | description |
|------|-------------|
| thread_id | __Required__ thread_id parameter |

## activity get-thread-subscription-for-authenticated-user

https://developer.github.com/v3/activity/notifications/#get-a-thread-subscription-for-the-authenticated-user

This checks to see if the current user is subscribed to a thread. You can also [get a repository subscription](https://developer.github.com/v3/activity/watching/#get-a-repository-subscription).

Note that subscriptions are only generated if a user is participating in a conversation--for example, they've replied to the thread, were **@mentioned**, or manually subscribe to a thread.

### parameters


| name | description |
|------|-------------|
| thread_id | __Required__ thread_id parameter |

## activity list-events-for-authenticated-user

https://developer.github.com/v3/activity/events/#list-events-for-the-authenticated-user

If you are authenticated as the given user, you will see your private events. Otherwise, you'll only see public events.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-notifications-for-authenticated-user

https://developer.github.com/v3/activity/notifications/#list-notifications-for-the-authenticated-user

List all notifications for the current user, sorted by most recently updated.

### parameters


| name | description |
|------|-------------|
| all | If `true`, show notifications marked as read. |
| before | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| page | Page number of the results to fetch. |
| participating | If `true`, only shows notifications in which the user is directly participating or mentioned. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## activity list-org-events-for-authenticated-user

https://developer.github.com/v3/activity/events/#list-organization-events-for-the-authenticated-user

This is the user's organization dashboard. You must be authenticated as the user to view this.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-public-events

https://developer.github.com/v3/activity/events/#list-public-events

We delay the public events feed by five minutes, which means the most recent event returned by the public events API actually occurred at least five minutes ago.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-public-events-for-repo-network

https://developer.github.com/v3/activity/events/#list-public-events-for-a-network-of-repositories



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-public-events-for-user

https://developer.github.com/v3/activity/events/#list-public-events-for-a-user



### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-public-org-events

https://developer.github.com/v3/activity/events/#list-public-organization-events



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-received-events-for-user

https://developer.github.com/v3/activity/events/#list-events-received-by-the-authenticated-user

These are events that you've received by watching repos and following users. If you are authenticated as the given user, you will see private events. Otherwise, you'll only see public events.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-received-public-events-for-user

https://developer.github.com/v3/activity/events/#list-public-events-received-by-a-user



### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-repo-events

https://developer.github.com/v3/activity/events/#list-repository-events



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-repo-notifications-for-authenticated-user

https://developer.github.com/v3/activity/notifications/#list-repository-notifications-for-the-authenticated-user

List all notifications for the current user.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| all | If `true`, show notifications marked as read. |
| before | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| page | Page number of the results to fetch. |
| participating | If `true`, only shows notifications in which the user is directly participating or mentioned. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## activity list-repos-starred-by-authenticated-user

https://developer.github.com/v3/activity/starring/#list-repositories-starred-by-the-authenticated-user

Lists repositories the authenticated user has starred.

You can also find out _when_ stars were created by passing the following custom [media type](https://developer.github.com/v3/media/) via the `Accept` header:

### parameters


| name | description |
|------|-------------|
| direction | One of `asc` (ascending) or `desc` (descending). |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | One of `created` (when the repository was starred) or `updated` (when it was last pushed to). |

## activity list-repos-starred-by-user

https://developer.github.com/v3/activity/starring/#list-repositories-starred-by-a-user

Lists repositories a user has starred.

You can also find out _when_ stars were created by passing the following custom [media type](https://developer.github.com/v3/media/) via the `Accept` header:

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| direction | One of `asc` (ascending) or `desc` (descending). |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | One of `created` (when the repository was starred) or `updated` (when it was last pushed to). |

## activity list-repos-watched-by-user

https://developer.github.com/v3/activity/watching/#list-repositories-watched-by-a-user

Lists repositories a user is watching.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-stargazers-for-repo

https://developer.github.com/v3/activity/starring/#list-stargazers

Lists the people that have starred the repository.

You can also find out _when_ stars were created by passing the following custom [media type](https://developer.github.com/v3/media/) via the `Accept` header:

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-watched-repos-for-authenticated-user

https://developer.github.com/v3/activity/watching/#list-repositories-watched-by-the-authenticated-user

Lists repositories the authenticated user is watching.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity list-watchers-for-repo

https://developer.github.com/v3/activity/watching/#list-watchers

Lists the people watching the specified repository.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## activity mark-notifications-as-read

https://developer.github.com/v3/activity/notifications/#mark-notifications-as-read

Marks all notifications as "read" removes it from the [default view on GitHub](https://github.com/notifications). If the number of notifications is too large to complete in one request, you will receive a `202 Accepted` status and GitHub will run an asynchronous process to mark notifications as "read." To check whether any "unread" notifications remain, you can use the [List notifications for the authenticated user](https://developer.github.com/v3/activity/notifications/#list-notifications-for-the-authenticated-user) endpoint and pass the query parameter `all=false`.

### parameters


| name | description |
|------|-------------|
| last_read_at | Describes the last point that notifications were checked. |
| read | Whether the notification has been read. |

## activity mark-repo-notifications-as-read

https://developer.github.com/v3/activity/notifications/#mark-repository-notifications-as-read

Marks all notifications in a repository as "read" removes them from the [default view on GitHub](https://github.com/notifications). If the number of notifications is too large to complete in one request, you will receive a `202 Accepted` status and GitHub will run an asynchronous process to mark notifications as "read." To check whether any "unread" notifications remain, you can use the [List repository notifications for the authenticated user](https://developer.github.com/v3/activity/notifications/#list-repository-notifications-for-the-authenticated-user) endpoint and pass the query parameter `all=false`.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| last_read_at | Describes the last point that notifications were checked. Anything updated since this time will not be marked as read. If you omit this parameter, all notifications are marked as read. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. Default: The current timestamp. |

## activity mark-thread-as-read

https://developer.github.com/v3/activity/notifications/#mark-a-thread-as-read



### parameters


| name | description |
|------|-------------|
| thread_id | __Required__ thread_id parameter |

## activity set-repo-subscription

https://developer.github.com/v3/activity/watching/#set-a-repository-subscription

If you would like to watch a repository, set `subscribed` to `true`. If you would like to ignore notifications made within a repository, set `ignored` to `true`. If you would like to stop watching a repository, [delete the repository's subscription](https://developer.github.com/v3/activity/watching/#delete-a-repository-subscription) completely.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| ignored | Determines if all notifications should be blocked from this repository. |
| subscribed | Determines if notifications should be received from this repository. |

## activity set-thread-subscription

https://developer.github.com/v3/activity/notifications/#set-a-thread-subscription

If you are watching a repository, you receive notifications for all threads by default. Use this endpoint to ignore future notifications for threads until you comment on the thread or get an **@mention**.

You can also use this endpoint to subscribe to threads that you are currently not receiving notifications for or to subscribed to threads that you have previously ignored.

Unsubscribing from a conversation in a repository that you are not watching is functionally equivalent to the [Delete a thread subscription](https://developer.github.com/v3/activity/notifications/#delete-a-thread-subscription) endpoint.

### parameters


| name | description |
|------|-------------|
| thread_id | __Required__ thread_id parameter |
| ignored | Whether to block all notifications from a thread. |

## activity star-repo-for-authenticated-user

https://developer.github.com/v3/activity/starring/#star-a-repository-for-the-authenticated-user

Note that you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## activity unstar-repo-for-authenticated-user

https://developer.github.com/v3/activity/starring/#unstar-a-repository-for-the-authenticated-user



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

# apps


## apps add-repo-to-installation

https://developer.github.com/v3/apps/installations/#add-a-repository-to-an-app-installation

Add a single repository to an installation. The authenticated user must have admin access to the repository.

You must use a personal access token (which you can create via the [command line](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) or the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization)) or [Basic Authentication](https://developer.github.com/v3/auth/#basic-authentication) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| installation_id | __Required__ installation_id parameter |
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| repository_id | __Required__ repository_id parameter |

## apps check-authorization

https://developer.github.com/v3/apps/oauth_applications/#check-an-authorization

**Deprecation Notice:** GitHub will replace and discontinue OAuth endpoints containing `access_token` in the path parameter. We are introducing new endpoints that allow you to securely manage tokens for OAuth Apps by using `access_token` as an input parameter. The OAuth Application API will be removed on May 5, 2021. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-app-endpoint/).

OAuth applications can use a special API method for checking OAuth token validity without exceeding the normal rate limits for failed login attempts. Authentication works differently with this particular endpoint. You must use [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication) when accessing this endpoint, using the OAuth application's `client_id` and `client_secret` as the username and password. Invalid tokens will return `404 NOT FOUND`.

### parameters


| name | description |
|------|-------------|
| access_token | __Required__  |
| client_id | __Required__  |

## apps check-token

https://developer.github.com/v3/apps/oauth_applications/#check-a-token

OAuth applications can use a special API method for checking OAuth token validity without exceeding the normal rate limits for failed login attempts. Authentication works differently with this particular endpoint. You must use [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication) to use this endpoint, where the username is the OAuth application `client_id` and the password is its `client_secret`. Invalid tokens will return `404 NOT FOUND`.

### parameters


| name | description |
|------|-------------|
| access_token | __Required__ The access_token of the OAuth application. |
| client_id | __Required__  |

## apps create-content-attachment

https://developer.github.com/v3/apps/installations/#create-a-content-attachment

Creates an attachment under a content reference URL in the body or comment of an issue or pull request. Use the `id` of the content reference from the [`content_reference` event](https://developer.github.com/webhooks/event-payloads/#content_reference) to create an attachment.

The app must create a content attachment within six hours of the content reference URL being posted. See "[Using content attachments](https://developer.github.com/apps/using-content-attachments/)" for details about content attachments.

You must use an [installation access token](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-an-installation) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The body of the attachment |
| content_reference_id | __Required__ content_reference_id parameter |
| corsair-preview | __Required__ To access the Content Attachments API during the preview period, you must set this flag. |
| title | __Required__ The title of the attachment |

## apps create-from-manifest

https://developer.github.com/v3/apps/#create-a-github-app-from-a-manifest

Use this endpoint to complete the handshake necessary when implementing the [GitHub App Manifest flow](https://developer.github.com/apps/building-github-apps/creating-github-apps-from-a-manifest/). When you create a GitHub App with the manifest flow, you receive a temporary `code` used to retrieve the GitHub App's `id`, `pem` (private key), and `webhook_secret`.

### parameters


| name | description |
|------|-------------|
| code | __Required__ code parameter |

## apps create-installation-access-token

https://developer.github.com/v3/apps/#create-an-installation-access-token-for-an-app

Creates an installation access token that enables a GitHub App to make authenticated API requests for the app's installation on an organization or individual account. Installation tokens expire one hour from the time you create them. Using an expired token produces a status code of `401 - Unauthorized`, and requires creating a new installation token. By default the installation token has access to all repositories that the installation can access. To restrict the access to specific repositories, you can provide the `repository_ids` when creating the token. When you omit `repository_ids`, the response does not contain the `repositories` key.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| installation_id | __Required__ installation_id parameter |
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| permissions.contents |  |
| permissions.def_not_a_repo |  |
| permissions.deployments |  |
| permissions.issues |  |
| permissions.single_file |  |
| repositories | List of repository names that the token should have access to |
| repository_ids | List of repository IDs that the token should have access to |

## apps delete-authorization

https://developer.github.com/v3/apps/oauth_applications/#delete-an-app-authorization

OAuth application owners can revoke a grant for their OAuth application and a specific user. You must use [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication) when accessing this endpoint, using the OAuth application's `client_id` and `client_secret` as the username and password. You must also provide a valid OAuth `access_token` as an input parameter and the grant for the token's owner will be deleted.
Deleting an OAuth application's grant will also delete all OAuth tokens associated with the application for the user. Once deleted, the application will have no access to the user's account and will no longer be listed on [the application authorizations settings screen within GitHub](https://github.com/settings/applications#authorized).

### parameters


| name | description |
|------|-------------|
| client_id | __Required__  |
| access_token | The OAuth access token used to authenticate to the GitHub API. |

## apps delete-installation

https://developer.github.com/v3/apps/#delete-an-installation-for-the-authenticated-app

Uninstalls a GitHub App on a user, organization, or business account. If you prefer to temporarily suspend an app's access to your account's resources, then we recommend the "[Suspend an app installation](https://developer.github.com/v3/apps/#suspend-an-app-installation)" endpoint.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| installation_id | __Required__ installation_id parameter |
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |

## apps delete-token

https://developer.github.com/v3/apps/oauth_applications/#delete-an-app-token

OAuth application owners can revoke a single token for an OAuth application. You must use [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication) when accessing this endpoint, using the OAuth application's `client_id` and `client_secret` as the username and password.

### parameters


| name | description |
|------|-------------|
| client_id | __Required__  |
| access_token | The OAuth access token used to authenticate to the GitHub API. |

## apps get-authenticated

https://developer.github.com/v3/apps/#get-the-authenticated-app

Returns the GitHub App associated with the authentication credentials used. To see how many app installations are associated with this GitHub App, see the `installations_count` in the response. For more details about your app's installations, see the "[List installations for the authenticated app](https://developer.github.com/v3/apps/#list-installations-for-the-authenticated-app)" endpoint.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |

## apps get-by-slug

https://developer.github.com/v3/apps/#get-an-app

**Note**: The `:app_slug` is just the URL-friendly name of your GitHub App. You can find this on the settings page for your GitHub App (e.g., `https://github.com/settings/apps/:app_slug`).

If the GitHub App you specify is public, you can access this endpoint without authenticating. If the GitHub App you specify is private, you must authenticate with a [personal access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) or an [installation access token](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-an-installation) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| app_slug | __Required__ app_slug parameter |
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |

## apps get-installation

https://developer.github.com/v3/apps/#get-an-installation-for-the-authenticated-app

Enables an authenticated GitHub App to find an installation's information using the installation id. The installation's account type (`target_type`) will be either an organization or a user account, depending which account the repository belongs to.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| installation_id | __Required__ installation_id parameter |
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |

## apps get-org-installation

https://developer.github.com/v3/apps/#get-an-organization-installation-for-the-authenticated-app

Enables an authenticated GitHub App to find the organization's installation information.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| org | __Required__  |

## apps get-repo-installation

https://developer.github.com/v3/apps/#get-a-repository-installation-for-the-authenticated-app

Enables an authenticated GitHub App to find the repository's installation information. The installation's account type will be either an organization or a user account, depending which account the repository belongs to.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| repo | __Required__ repository in OWNER/REPO form |

## apps get-subscription-plan-for-account

https://developer.github.com/v3/apps/marketplace/#get-a-subscription-plan-for-an-account

Shows whether the user or organization account actively subscribes to a plan listed by the authenticated GitHub App. When someone submits a plan change that won't be processed until the end of their billing cycle, you will also see the upcoming pending change.

GitHub Apps must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth Apps must use [basic authentication](https://developer.github.com/v3/auth/#basic-authentication) with their client ID and client secret to access this endpoint.

### parameters


| name | description |
|------|-------------|
| account_id | __Required__ account_id parameter |

## apps get-subscription-plan-for-account-stubbed

https://developer.github.com/v3/apps/marketplace/#get-a-subscription-plan-for-an-account-stubbed

Shows whether the user or organization account actively subscribes to a plan listed by the authenticated GitHub App. When someone submits a plan change that won't be processed until the end of their billing cycle, you will also see the upcoming pending change.

GitHub Apps must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth Apps must use [basic authentication](https://developer.github.com/v3/auth/#basic-authentication) with their client ID and client secret to access this endpoint.

### parameters


| name | description |
|------|-------------|
| account_id | __Required__ account_id parameter |

## apps get-user-installation

https://developer.github.com/v3/apps/#get-a-user-installation-for-the-authenticated-app

Enables an authenticated GitHub App to find the user’s installation information.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| username | __Required__  |

## apps list-accounts-for-plan

https://developer.github.com/v3/apps/marketplace/#list-accounts-for-a-plan

Returns user and organization accounts associated with the specified plan, including free plans. For per-seat pricing, you see the list of accounts that have purchased the plan, including the number of seats purchased. When someone submits a plan change that won't be processed until the end of their billing cycle, you will also see the upcoming pending change.

GitHub Apps must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth Apps must use [basic authentication](https://developer.github.com/v3/auth/#basic-authentication) with their client ID and client secret to access this endpoint.

### parameters


| name | description |
|------|-------------|
| plan_id | __Required__ plan_id parameter |
| direction | To return the oldest accounts first, set to `asc`. Can be one of `asc` or `desc`. Ignored without the `sort` parameter. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | One of `created` (when the repository was starred) or `updated` (when it was last pushed to). |

## apps list-accounts-for-plan-stubbed

https://developer.github.com/v3/apps/marketplace/#list-accounts-for-a-plan-stubbed

Returns repository and organization accounts associated with the specified plan, including free plans. For per-seat pricing, you see the list of accounts that have purchased the plan, including the number of seats purchased. When someone submits a plan change that won't be processed until the end of their billing cycle, you will also see the upcoming pending change.

GitHub Apps must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth Apps must use [basic authentication](https://developer.github.com/v3/auth/#basic-authentication) with their client ID and client secret to access this endpoint.

### parameters


| name | description |
|------|-------------|
| plan_id | __Required__ plan_id parameter |
| direction | To return the oldest accounts first, set to `asc`. Can be one of `asc` or `desc`. Ignored without the `sort` parameter. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | One of `created` (when the repository was starred) or `updated` (when it was last pushed to). |

## apps list-installation-repos-for-authenticated-user

https://developer.github.com/v3/apps/installations/#list-repositories-accessible-to-the-user-access-token

List repositories that the authenticated user has explicit permission (`:read`, `:write`, or `:admin`) to access for an installation.

The authenticated user has explicit permission to access repositories they own, repositories where they are a collaborator, and repositories that they can access through an organization membership.

You must use a [user-to-server OAuth access token](https://developer.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/#identifying-users-on-your-site), created for a user who has authorized your GitHub App, to access this endpoint.

The access the user has to each repository is included in the hash under the `permissions` key.

### parameters


| name | description |
|------|-------------|
| installation_id | __Required__ installation_id parameter |
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| mercy-preview | The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## apps list-installations

https://developer.github.com/v3/apps/#list-installations-for-the-authenticated-app

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

The permissions the installation has are included under the `permissions` key.

### parameters


| name | description |
|------|-------------|
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| outdated |  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## apps list-installations-for-authenticated-user

https://developer.github.com/v3/apps/installations/#list-app-installations-accessible-to-the-user-access-token

Lists installations of your GitHub App that the authenticated user has explicit permission (`:read`, `:write`, or `:admin`) to access.

You must use a [user-to-server OAuth access token](https://developer.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/#identifying-users-on-your-site), created for a user who has authorized your GitHub App, to access this endpoint.

The authenticated user has explicit permission to access repositories they own, repositories where they are a collaborator, and repositories that they can access through an organization membership.

You can find the permissions for the installation under the `permissions` key.

### parameters


| name | description |
|------|-------------|
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## apps list-plans

https://developer.github.com/v3/apps/marketplace/#list-plans

Lists all plans that are part of your GitHub Marketplace listing.

GitHub Apps must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth Apps must use [basic authentication](https://developer.github.com/v3/auth/#basic-authentication) with their client ID and client secret to access this endpoint.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## apps list-plans-stubbed

https://developer.github.com/v3/apps/marketplace/#list-plans-stubbed

Lists all plans that are part of your GitHub Marketplace listing.

GitHub Apps must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth Apps must use [basic authentication](https://developer.github.com/v3/auth/#basic-authentication) with their client ID and client secret to access this endpoint.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## apps list-repos-accessible-to-installation

https://developer.github.com/v3/apps/installations/#list-repositories-accessible-to-the-app-installation

List repositories that an app installation can access.

You must use an [installation access token](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-an-installation) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| mercy-preview | The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## apps list-subscriptions-for-authenticated-user

https://developer.github.com/v3/apps/marketplace/#list-subscriptions-for-the-authenticated-user

Lists the active subscriptions for the authenticated user. You must use a [user-to-server OAuth access token](https://developer.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/#identifying-users-on-your-site), created for a user who has authorized your GitHub App, to access this endpoint. . OAuth Apps must authenticate using an [OAuth token](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/).

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## apps list-subscriptions-for-authenticated-user-stubbed

https://developer.github.com/v3/apps/marketplace/#list-subscriptions-for-the-authenticated-user-stubbed

Lists the active subscriptions for the authenticated user. You must use a [user-to-server OAuth access token](https://developer.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/#identifying-users-on-your-site), created for a user who has authorized your GitHub App, to access this endpoint. . OAuth Apps must authenticate using an [OAuth token](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/).

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## apps remove-repo-from-installation

https://developer.github.com/v3/apps/installations/#remove-a-repository-from-an-app-installation

Remove a single repository from an installation. The authenticated user must have admin access to the repository.

You must use a personal access token (which you can create via the [command line](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) or the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization)) or [Basic Authentication](https://developer.github.com/v3/auth/#basic-authentication) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| installation_id | __Required__ installation_id parameter |
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| repository_id | __Required__ repository_id parameter |

## apps reset-authorization

https://developer.github.com/v3/apps/oauth_applications/#reset-an-authorization

**Deprecation Notice:** GitHub will replace and discontinue OAuth endpoints containing `access_token` in the path parameter. We are introducing new endpoints that allow you to securely manage tokens for OAuth Apps by using `access_token` as an input parameter. The OAuth Application API will be removed on May 5, 2021. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-app-endpoint/).

OAuth applications can use this API method to reset a valid OAuth token without end-user involvement. Applications must save the "token" property in the response because changes take effect immediately. You must use [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication) when accessing this endpoint, using the OAuth application's `client_id` and `client_secret` as the username and password. Invalid tokens will return `404 NOT FOUND`.

### parameters


| name | description |
|------|-------------|
| access_token | __Required__  |
| client_id | __Required__  |

## apps reset-token

https://developer.github.com/v3/apps/oauth_applications/#reset-a-token

OAuth applications can use this API method to reset a valid OAuth token without end-user involvement. Applications must save the "token" property in the response because changes take effect immediately. You must use [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication) when accessing this endpoint, using the OAuth application's `client_id` and `client_secret` as the username and password. Invalid tokens will return `404 NOT FOUND`.

### parameters


| name | description |
|------|-------------|
| access_token | __Required__ The access_token of the OAuth application. |
| client_id | __Required__  |

## apps revoke-authorization-for-application

https://developer.github.com/v3/apps/oauth_applications/#revoke-an-authorization-for-an-application

**Deprecation Notice:** GitHub will replace and discontinue OAuth endpoints containing `access_token` in the path parameter. We are introducing new endpoints that allow you to securely manage tokens for OAuth Apps by using `access_token` as an input parameter. The OAuth Application API will be removed on May 5, 2021. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-app-endpoint/).

OAuth application owners can revoke a single token for an OAuth application. You must use [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication) when accessing this endpoint, using the OAuth application's `client_id` and `client_secret` as the username and password.

### parameters


| name | description |
|------|-------------|
| access_token | __Required__  |
| client_id | __Required__  |

## apps revoke-grant-for-application

https://developer.github.com/v3/apps/oauth_applications/#revoke-a-grant-for-an-application

**Deprecation Notice:** GitHub will replace and discontinue OAuth endpoints containing `access_token` in the path parameter. We are introducing new endpoints that allow you to securely manage tokens for OAuth Apps by using `access_token` as an input parameter. The OAuth Application API will be removed on May 5, 2021. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-app-endpoint/).

OAuth application owners can revoke a grant for their OAuth application and a specific user. You must use [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication) when accessing this endpoint, using the OAuth application's `client_id` and `client_secret` as the username and password. You must also provide a valid token as `:access_token` and the grant for the token's owner will be deleted.

Deleting an OAuth application's grant will also delete all OAuth tokens associated with the application for the user. Once deleted, the application will have no access to the user's account and will no longer be listed on [the Applications settings page under "Authorized OAuth Apps" on GitHub](https://github.com/settings/applications#authorized).

### parameters


| name | description |
|------|-------------|
| access_token | __Required__  |
| client_id | __Required__  |

## apps revoke-installation-access-token

https://developer.github.com/v3/apps/installations/#revoke-an-installation-access-token

Revokes the installation token you're using to authenticate as an installation and access this endpoint.

Once an installation token is revoked, the token is invalidated and cannot be used. Other endpoints that require the revoked installation token must have a new installation token to work. You can create a new token using the "[Create an installation access token for an app](https://developer.github.com/v3/apps/#create-an-installation-access-token-for-an-app)" endpoint.

You must use an [installation access token](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-an-installation) to access this endpoint.

## apps suspend-installation

https://developer.github.com/v3/apps/#suspend-an-app-installation

**Note:** Suspending a GitHub App installation is currently in beta and subject to change. Before you can suspend a GitHub App, the app owner must enable suspending installations for the app by opting-in to the beta. For more information, see "[Suspending a GitHub App installation](https://developer.github.com/apps/managing-github-apps/suspending-a-github-app-installation/)."

Suspends a GitHub App on a user, organization, or business account, which blocks the app from accessing the account's resources. When a GitHub App is suspended, the app's access to the GitHub API or webhook events is blocked for that account.

To suspend a GitHub App, you must be an account owner or have admin permissions in the repository or organization where the app is installed.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| installation_id | __Required__ installation_id parameter |

## apps unsuspend-installation

https://developer.github.com/v3/apps/#unsuspend-an-app-installation

**Note:** Suspending a GitHub App installation is currently in beta and subject to change. Before you can suspend a GitHub App, the app owner must enable suspending installations for the app by opting-in to the beta. For more information, see "[Suspending a GitHub App installation](https://developer.github.com/apps/managing-github-apps/suspending-a-github-app-installation/)."

Removes a GitHub App installation suspension.

To unsuspend a GitHub App, you must be an account owner or have admin permissions in the repository or organization where the app is installed and suspended.

You must use a [JWT](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.

### parameters


| name | description |
|------|-------------|
| installation_id | __Required__ installation_id parameter |

# billing


## billing get-github-actions-billing-ghe

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-enterprise

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the summary of the free and paid GitHub Actions minutes used.

Paid minutes only apply to workflows in private repositories that use GitHub-hosted runners. Minutes used is listed for each GitHub-hosted runner operating system. Any job re-runs are also included in the usage. The usage does not include the multiplier for macOS and Windows runners and is not rounded up to the nearest whole minute. For more information, see "[Managing billing for GitHub Actions](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-actions)".

The authenticated user must be an enterprise admin.

### parameters


| name | description |
|------|-------------|
| enterprise_id | __Required__ Unique identifier of the GitHub Enterprise Cloud instance. |

## billing get-github-actions-billing-org

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-organization

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the summary of the free and paid GitHub Actions minutes used.

Paid minutes only apply to workflows in private repositories that use GitHub-hosted runners. Minutes used is listed for each GitHub-hosted runner operating system. Any job re-runs are also included in the usage. The usage does not include the multiplier for macOS and Windows runners and is not rounded up to the nearest whole minute. For more information, see "[Managing billing for GitHub Actions](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-actions)".

Access tokens must have the `read:org` scope.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## billing get-github-actions-billing-user

https://developer.github.com/v3/billing/#get-github-actions-billing-for-a-user

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the summary of the free and paid GitHub Actions minutes used.

Paid minutes only apply to workflows in private repositories that use GitHub-hosted runners. Minutes used is listed for each GitHub-hosted runner operating system. Any job re-runs are also included in the usage. The usage does not include the multiplier for macOS and Windows runners and is not rounded up to the nearest whole minute. For more information, see "[Managing billing for GitHub Actions](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-actions)".

Access tokens must have the `user` scope.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## billing get-github-packages-billing-ghe

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-enterprise

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the free and paid storage used for GitHub Packages in gigabytes.

Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."

The authenticated user must be an enterprise admin.

### parameters


| name | description |
|------|-------------|
| enterprise_id | __Required__ Unique identifier of the GitHub Enterprise Cloud instance. |

## billing get-github-packages-billing-org

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-organization

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the free and paid storage usued for GitHub Packages in gigabytes.

Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."

Access tokens must have the `read:org` scope.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## billing get-github-packages-billing-user

https://developer.github.com/v3/billing/#get-github-packages-billing-for-a-user

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the free and paid storage used for GitHub Packages in gigabytes.

Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."

Access tokens must have the `user` scope.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## billing get-shared-storage-billing-ghe

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-enterprise

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the estimated paid and estimated total storage used for GitHub Actions and Github Packages.

Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."

The authenticated user must be an enterprise admin.

### parameters


| name | description |
|------|-------------|
| enterprise_id | __Required__ Unique identifier of the GitHub Enterprise Cloud instance. |

## billing get-shared-storage-billing-org

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-organization

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the estimated paid and estimated total storage used for GitHub Actions and Github Packages.

Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."

Access tokens must have the `read:org` scope.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## billing get-shared-storage-billing-user

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-a-user

**Warning:** The Billing API is currently in public beta and subject to change.

Gets the estimated paid and estimated total storage used for GitHub Actions and Github Packages.

Paid minutes only apply to packages stored for private repositories. For more information, see "[Managing billing for GitHub Packages](https://help.github.com/github/setting-up-and-managing-billing-and-payments-on-github/managing-billing-for-github-packages)."

Access tokens must have the `user` scope.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |

# checks


## checks create

https://developer.github.com/v3/checks/runs/#create-a-check-run

**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.

Creates a new check run for a specific commit in a repository. Your GitHub App must have the `checks:write` permission to create check runs.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| head_sha | __Required__ The SHA of the commit. |
| name | __Required__ The name of the check. For example, "code-coverage". |
| repo | __Required__ repository in OWNER/REPO form |
| actions | Displays a button on GitHub that can be clicked to alert your app to do additional tasks. For example, a code linting app can display a button that automatically fixes detected errors. The button created in this object is displayed after the check run completes. When a user clicks the button, GitHub sends the [`check_run.requested_action` webhook](https://developer.github.com/webhooks/event-payloads/#check_run) to your app. Each action includes a `label`, `identifier` and `description`. A maximum of three actions are accepted. See the [`actions` object](https://developer.github.com/v3/checks/runs/#actions-object) description. To learn more about check runs and requested actions, see "[Check runs and requested actions](https://developer.github.com/v3/checks/runs/#check-runs-and-requested-actions)." To learn more about check runs and requested actions, see "[Check runs and requested actions](https://developer.github.com/v3/checks/runs/#check-runs-and-requested-actions)." |
| completed_at | The time the check completed. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| conclusion | **Required if you provide `completed_at` or a `status` of `completed`**. The final conclusion of the check. Can be one of `success`, `failure`, `neutral`, `cancelled`, `skipped`, `timed_out`, or `action_required`. When the conclusion is `action_required`, additional details should be provided on the site specified by `details_url`.  <br>**Note:** Providing `conclusion` will automatically set the `status` parameter to `completed`. Only GitHub can change a check run conclusion to `stale`. |
| details_url | The URL of the integrator's site that has the full details of the check. If the integrator does not provide this, then the homepage of the GitHub app is used. |
| external_id | A reference for the run on the integrator's system. |
| output.annotations | Adds information from your analysis to specific lines of code. Annotations are visible on GitHub in the **Checks** and **Files changed** tab of the pull request. The Checks API limits the number of annotations to a maximum of 50 per API request. To create more than 50 annotations, you have to make multiple requests to the [Update a check run](https://developer.github.com/v3/checks/runs/#update-a-check-run) endpoint. Each time you update the check run, annotations are appended to the list of annotations that already exist for the check run. For details about how you can view annotations on GitHub, see "[About status checks](https://help.github.com/articles/about-status-checks#checks)". See the [`annotations` object](https://developer.github.com/v3/checks/runs/#annotations-object) description for details about how to use this parameter. |
| output.images | Adds images to the output displayed in the GitHub pull request UI. See the [`images` object](https://developer.github.com/v3/checks/runs/#images-object) description for details. |
| output.summary | The summary of the check run. This parameter supports Markdown. |
| output.text | The details of the check run. This parameter supports Markdown. |
| output.title | The title of the check run. |
| started_at | The time that the check run began. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| status | The current status. Can be one of `queued`, `in_progress`, or `completed`. |

## checks create-suite

https://developer.github.com/v3/checks/suites/#create-a-check-suite

**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.

By default, check suites are automatically created when you create a [check run](https://developer.github.com/v3/checks/runs/). You only need to use this endpoint for manually creating check suites when you've disabled automatic creation using "[Update repository preferences for check suites](https://developer.github.com/v3/checks/suites/#update-repository-preferences-for-check-suites)". Your GitHub App must have the `checks:write` permission to create check suites.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| head_sha | __Required__ The sha of the head commit. |
| repo | __Required__ repository in OWNER/REPO form |

## checks get

https://developer.github.com/v3/checks/runs/#get-a-check-run

**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.

Gets a single check run using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth Apps and authenticated users must have the `repo` scope to get check runs in a private repository.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| check_run_id | __Required__ check_run_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## checks get-suite

https://developer.github.com/v3/checks/suites/#get-a-check-suite

**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.

Gets a single check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check suites. OAuth Apps and authenticated users must have the `repo` scope to get check suites in a private repository.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| check_suite_id | __Required__ check_suite_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## checks list-annotations

https://developer.github.com/v3/checks/runs/#list-check-run-annotations

Lists annotations for a check run using the annotation `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get annotations for a check run. OAuth Apps and authenticated users must have the `repo` scope to get annotations for a check run in a private repository.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| check_run_id | __Required__ check_run_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## checks list-for-ref

https://developer.github.com/v3/checks/runs/#list-check-runs-for-a-git-reference

**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.

Lists check runs for a commit ref. The `ref` can be a SHA, branch name, or a tag name. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth Apps and authenticated users must have the `repo` scope to get check runs in a private repository.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| check_name | Returns check runs with the specified `name`. |
| filter | Filters check runs by their `completed_at` timestamp. Can be one of `latest` (returning the most recent check runs) or `all`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| status | Returns check runs with the specified `status`. Can be one of `queued`, `in_progress`, or `completed`. |

## checks list-for-suite

https://developer.github.com/v3/checks/runs/#list-check-runs-in-a-check-suite

**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.

Lists check runs for a check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth Apps and authenticated users must have the `repo` scope to get check runs in a private repository.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| check_suite_id | __Required__ check_suite_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| check_name | Returns check runs with the specified `name`. |
| filter | Filters check runs by their `completed_at` timestamp. Can be one of `latest` (returning the most recent check runs) or `all`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| status | Returns check runs with the specified `status`. Can be one of `queued`, `in_progress`, or `completed`. |

## checks list-suites-for-ref

https://developer.github.com/v3/checks/suites/#list-check-suites-for-a-git-reference

**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.

Lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to list check suites. OAuth Apps and authenticated users must have the `repo` scope to get check suites in a private repository.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| app_id | Filters check suites by GitHub App `id`. |
| check_name | Returns check runs with the specified `name`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## checks rerequest-suite

https://developer.github.com/v3/checks/suites/#rerequest-a-check-suite

Triggers GitHub to rerequest an existing check suite, without pushing new code to a repository. This endpoint will trigger the [`check_suite` webhook](https://developer.github.com/webhooks/event-payloads/#check_suite) event with the action `rerequested`. When a check suite is `rerequested`, its `status` is reset to `queued` and the `conclusion` is cleared.

To rerequest a check suite, your GitHub App must have the `checks:read` permission on a private repository or pull access to a public repository.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| check_suite_id | __Required__ check_suite_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## checks set-suites-preferences

https://developer.github.com/v3/checks/suites/#update-repository-preferences-for-check-suites

Changes the default automatic flow when creating check suites. By default, a check suite is automatically created each time code is pushed to a repository. When you disable the automatic creation of check suites, you can manually [Create a check suite](https://developer.github.com/v3/checks/suites/#create-a-check-suite). You must have admin permissions in the repository to set preferences for check suites.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |
| auto_trigger_checks | Enables or disables automatic creation of CheckSuite events upon pushes to the repository. Enabled by default. See the [`auto_trigger_checks` object](https://developer.github.com/v3/checks/suites/#auto_trigger_checks-object) description for details. |

## checks update

https://developer.github.com/v3/checks/runs/#update-a-check-run

**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.

Updates a check run for a specific commit in a repository. Your GitHub App must have the `checks:write` permission to edit check runs.

### parameters


| name | description |
|------|-------------|
| antiope-preview | __Required__ The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must set this flag. |
| check_run_id | __Required__ check_run_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| actions | Possible further actions the integrator can perform, which a user may trigger. Each action includes a `label`, `identifier` and `description`. A maximum of three actions are accepted. See the [`actions` object](https://developer.github.com/v3/checks/runs/#actions-object) description. To learn more about check runs and requested actions, see "[Check runs and requested actions](https://developer.github.com/v3/checks/runs/#check-runs-and-requested-actions)." |
| completed_at | The time the check completed. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| conclusion | **Required if you provide `completed_at` or a `status` of `completed`**. The final conclusion of the check. Can be one of `success`, `failure`, `neutral`, `cancelled`, `skipped`, `timed_out`, or `action_required`.  <br>**Note:** Providing `conclusion` will automatically set the `status` parameter to `completed`. Only GitHub can change a check run conclusion to `stale`. |
| details_url | The URL of the integrator's site that has the full details of the check. |
| external_id | A reference for the run on the integrator's system. |
| name | The name of the check. For example, "code-coverage". |
| output.annotations | Adds information from your analysis to specific lines of code. Annotations are visible in GitHub's pull request UI. Annotations are visible in GitHub's pull request UI. The Checks API limits the number of annotations to a maximum of 50 per API request. To create more than 50 annotations, you have to make multiple requests to the [Update a check run](https://developer.github.com/v3/checks/runs/#update-a-check-run) endpoint. Each time you update the check run, annotations are appended to the list of annotations that already exist for the check run. For details about annotations in the UI, see "[About status checks](https://help.github.com/articles/about-status-checks#checks)". See the [`annotations` object](https://developer.github.com/v3/checks/runs/#annotations-object-1) description for details. |
| output.images | Adds images to the output displayed in the GitHub pull request UI. See the [`images` object](https://developer.github.com/v3/checks/runs/#annotations-object-1) description for details. |
| output.summary | Can contain Markdown. |
| output.text | Can contain Markdown. |
| output.title | **Required**. |
| started_at | This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| status | The current status. Can be one of `queued`, `in_progress`, or `completed`. |

# code-scanning


## code-scanning get-alert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert

Gets a single code scanning alert. You must use an access token with the `security_events` scope to use this endpoint. GitHub Apps must have the `security_events` read permission to use this endpoint.

The security `alert_id` is found at the end of the security alert's URL. For example, the security alert ID for `https://github.com/Octo-org/octo-repo/security/code-scanning/88` is `88`.

### parameters


| name | description |
|------|-------------|
| alert_id | __Required__ alert_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## code-scanning list-alerts-for-repo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository

Lists all open code scanning alerts for the default branch (usually `master`) and protected branches in a repository. You must use an access token with the `security_events` scope to use this endpoint. GitHub Apps must have the `security_events` read permission to use this endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| ref | Returns a list of code scanning alerts for a specific brach reference. The `ref` must be formatted as `heads/<branch name>`. |
| state | Set to `closed` to list only closed code scanning alerts. |

# codes-of-conduct


## codes-of-conduct get-all-codes-of-conduct

https://developer.github.com/v3/codes_of_conduct/#get-all-codes-of-conduct



### parameters


| name | description |
|------|-------------|
| scarlet-witch-preview | __Required__ The Codes of Conduct API is currently available for developers to preview.<br><br>To access the API during the preview period, you must set this flag. |

## codes-of-conduct get-conduct-code

https://developer.github.com/v3/codes_of_conduct/#get-a-code-of-conduct



### parameters


| name | description |
|------|-------------|
| key | __Required__ key parameter |
| scarlet-witch-preview | __Required__ The Codes of Conduct API is currently available for developers to preview.<br><br>To access the API during the preview period, you must set this flag. |

## codes-of-conduct get-for-repo

https://developer.github.com/v3/codes_of_conduct/#get-the-code-of-conduct-for-a-repository

This method returns the contents of the repository's code of conduct file, if one is detected.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| scarlet-witch-preview | __Required__ The Codes of Conduct API is currently available for developers to preview.<br><br>To access the API during the preview period, you must set this flag. |

# emojis


## emojis get

https://developer.github.com/v3/emojis/#get-emojis

Lists all the emojis available to use on GitHub.

# gists


## gists check-is-starred

https://developer.github.com/v3/gists/#check-if-a-gist-is-starred



### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |

## gists create

https://developer.github.com/v3/gists/#create-a-gist

Allows you to add a new gist with one or more files.

**Note:** Don't name your files "gistfile" with a numerical suffix. This is the format of the automatic naming scheme that Gist uses internally.

### parameters


| name | description |
|------|-------------|
| files | __Required__ Names and content for the files that make up the gist |
| description | Description of the gist |
| public |  |

## gists create-comment

https://developer.github.com/v3/gists/comments/#create-a-gist-comment



### parameters


| name | description |
|------|-------------|
| body | __Required__ The comment text. |
| gist_id | __Required__ gist_id parameter |

## gists delete

https://developer.github.com/v3/gists/#delete-a-gist



### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |

## gists delete-comment

https://developer.github.com/v3/gists/comments/#delete-a-gist-comment



### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| gist_id | __Required__ gist_id parameter |

## gists fork

https://developer.github.com/v3/gists/#fork-a-gist

**Note**: This was previously `/gists/:gist_id/fork`.

### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |

## gists get

https://developer.github.com/v3/gists/#get-a-gist



### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |

## gists get-comment

https://developer.github.com/v3/gists/comments/#get-a-gist-comment



### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| gist_id | __Required__ gist_id parameter |

## gists get-revision

https://developer.github.com/v3/gists/#get-a-gist-revision



### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |
| sha | __Required__ sha parameter |

## gists list

https://developer.github.com/v3/gists/#list-gists-for-the-authenticated-user

Lists the authenticated user's gists or if called anonymously, this endpoint returns all public gists:

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## gists list-comments

https://developer.github.com/v3/gists/comments/#list-gist-comments



### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## gists list-commits

https://developer.github.com/v3/gists/#list-gist-commits



### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## gists list-for-user

https://developer.github.com/v3/gists/#list-gists-for-a-user

Lists public gists for the specified user:

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## gists list-forks

https://developer.github.com/v3/gists/#list-gist-forks



### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## gists list-public

https://developer.github.com/v3/gists/#list-public-gists

List public gists sorted by most recently updated to least recently updated.

Note: With [pagination](https://developer.github.com/v3/#pagination), you can fetch up to 3000 gists. For example, you can fetch 100 pages with 30 gists per page or 30 pages with 100 gists per page.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## gists list-starred

https://developer.github.com/v3/gists/#list-starred-gists

List the authenticated user's starred gists:

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## gists star

https://developer.github.com/v3/gists/#star-a-gist

Note that you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |

## gists unstar

https://developer.github.com/v3/gists/#unstar-a-gist



### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |

## gists update

https://developer.github.com/v3/gists/#update-a-gist

Allows you to update or delete a gist file and rename gist files. Files from the previous version of the gist that aren't explicitly changed during an edit are unchanged.

### parameters


| name | description |
|------|-------------|
| gist_id | __Required__ gist_id parameter |
| description | Description of the gist |
| files | Names of files to be updated |

## gists update-comment

https://developer.github.com/v3/gists/comments/#update-a-gist-comment



### parameters


| name | description |
|------|-------------|
| body | __Required__ The comment text. |
| comment_id | __Required__ comment_id parameter |
| gist_id | __Required__ gist_id parameter |

# git


## git create-blob

https://developer.github.com/v3/git/blobs/#create-a-blob



### parameters


| name | description |
|------|-------------|
| content | __Required__ The new blob's content. |
| repo | __Required__ repository in OWNER/REPO form |
| encoding | The encoding used for `content`. Currently, `"utf-8"` and `"base64"` are supported. |

## git create-commit

https://developer.github.com/v3/git/commits/#create-a-commit

Creates a new Git [commit object](https://git-scm.com/book/en/v1/Git-Internals-Git-Objects#Commit-Objects).

In this example, the payload of the signature would be:



**Signature verification object**

The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:

These are the possible values for `reason` in the `verification` object:

| Value                    | Description                                                                                                                       |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| `expired_key`            | The key that made the signature is expired.                                                                                       |
| `not_signing_key`        | The "signing" flag is not among the usage flags in the GPG key that made the signature.                                           |
| `gpgverify_error`        | There was an error communicating with the signature verification service.                                                         |
| `gpgverify_unavailable`  | The signature verification service is currently unavailable.                                                                      |
| `unsigned`               | The object does not include a signature.                                                                                          |
| `unknown_signature_type` | A non-PGP signature was found in the commit.                                                                                      |
| `no_user`                | No user was associated with the `committer` email address in the commit.                                                          |
| `unverified_email`       | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. |
| `bad_email`              | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature.             |
| `unknown_key`            | The key that made the signature has not been registered with any user's account.                                                  |
| `malformed_signature`    | There was an error parsing the signature.                                                                                         |
| `invalid`                | The signature could not be cryptographically verified using the key whose key-id was found in the signature.                      |
| `valid`                  | None of the above errors applied, so the signature is considered to be verified.                                                  |

### parameters


| name | description |
|------|-------------|
| message | __Required__ The commit message |
| repo | __Required__ repository in OWNER/REPO form |
| tree | __Required__ The SHA of the tree object this commit points to |
| author.date | Indicates when this commit was authored (or committed). This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| author.email | The email of the author (or committer) of the commit |
| author.name | The name of the author (or committer) of the commit |
| committer.date | Indicates when this commit was authored (or committed). This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| committer.email | The email of the author (or committer) of the commit |
| committer.name | The name of the author (or committer) of the commit |
| parents | The SHAs of the commits that were the parents of this commit. If omitted or empty, the commit will be written as a root commit. For a single parent, an array of one SHA should be provided; for a merge commit, an array of more than one should be provided. |
| signature | The [PGP signature](https://en.wikipedia.org/wiki/Pretty_Good_Privacy) of the commit. GitHub adds the signature to the `gpgsig` header of the created commit. For a commit signature to be verifiable by Git or GitHub, it must be an ASCII-armored detached PGP signature over the string commit as it would be written to the object database. To pass a `signature` parameter, you need to first manually create a valid PGP signature, which can be complicated. You may find it easier to [use the command line](https://git-scm.com/book/id/v2/Git-Tools-Signing-Your-Work) to create signed commits. |

## git create-ref

https://developer.github.com/v3/git/refs/#create-a-reference

Creates a reference for your repository. You are unable to create new references for empty repositories, even if the commit SHA-1 hash used exists. Empty repositories are repositories without branches.

### parameters


| name | description |
|------|-------------|
| ref | __Required__ The name of the fully qualified reference (ie: `refs/heads/master`). If it doesn't start with 'refs' and have at least two slashes, it will be rejected. |
| repo | __Required__ repository in OWNER/REPO form |
| sha | __Required__ The SHA1 value for this reference. |
| key |  |

## git create-tag

https://developer.github.com/v3/git/tags/#create-a-tag-object

Note that creating a tag object does not create the reference that makes a tag in Git. If you want to create an annotated tag in Git, you have to do this call to create the tag object, and then [create](https://developer.github.com/v3/git/refs/#create-a-reference) the `refs/tags/[tag]` reference. If you want to create a lightweight tag, you only have to [create](https://developer.github.com/v3/git/refs/#create-a-reference) the tag reference - this call would be unnecessary.

**Signature verification object**

The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:

These are the possible values for `reason` in the `verification` object:

| Value                    | Description                                                                                                                       |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| `expired_key`            | The key that made the signature is expired.                                                                                       |
| `not_signing_key`        | The "signing" flag is not among the usage flags in the GPG key that made the signature.                                           |
| `gpgverify_error`        | There was an error communicating with the signature verification service.                                                         |
| `gpgverify_unavailable`  | The signature verification service is currently unavailable.                                                                      |
| `unsigned`               | The object does not include a signature.                                                                                          |
| `unknown_signature_type` | A non-PGP signature was found in the commit.                                                                                      |
| `no_user`                | No user was associated with the `committer` email address in the commit.                                                          |
| `unverified_email`       | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. |
| `bad_email`              | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature.             |
| `unknown_key`            | The key that made the signature has not been registered with any user's account.                                                  |
| `malformed_signature`    | There was an error parsing the signature.                                                                                         |
| `invalid`                | The signature could not be cryptographically verified using the key whose key-id was found in the signature.                      |
| `valid`                  | None of the above errors applied, so the signature is considered to be verified.                                                  |

### parameters


| name | description |
|------|-------------|
| message | __Required__ The tag message. |
| object | __Required__ The SHA of the git object this is tagging. |
| repo | __Required__ repository in OWNER/REPO form |
| tag | __Required__ The tag's name. This is typically a version (e.g., "v0.0.1"). |
| type | __Required__ The type of the object we're tagging. Normally this is a `commit` but it can also be a `tree` or a `blob`. |
| tagger.date | When this object was tagged. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| tagger.email | The email of the author of the tag |
| tagger.name | The name of the author of the tag |

## git create-tree

https://developer.github.com/v3/git/trees/#create-a-tree

The tree creation API accepts nested entries. If you specify both a tree and a nested path modifying that tree, this endpoint will overwrite the contents of the tree with the new path contents, and create a new tree structure.

If you use this endpoint to add, delete, or modify the file contents in a tree, you will need to commit the tree and then update a branch to point to the commit. For more information see "[Create a commit](https://developer.github.com/v3/git/commits/#create-a-commit)" and "[Update a reference](https://developer.github.com/v3/git/refs/#update-a-reference)."

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| tree | __Required__ Objects (of `path`, `mode`, `type`, and `sha`) specifying a tree structure. |
| base_tree | The SHA1 of the tree you want to update with new data. If you don't set this, the commit will be created on top of everything; however, it will only contain your change, the rest of your files will show up as deleted. |

## git delete-ref

https://developer.github.com/v3/git/refs/#delete-a-reference



### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## git get-blob

https://developer.github.com/v3/git/blobs/#get-a-blob

The `content` in the response will always be Base64 encoded.

_Note_: This API supports blobs up to 100 megabytes in size.

### parameters


| name | description |
|------|-------------|
| file_sha | __Required__ file_sha parameter |
| repo | __Required__ repository in OWNER/REPO form |

## git get-commit

https://developer.github.com/v3/git/commits/#get-a-commit

Gets a Git [commit object](https://git-scm.com/book/en/v1/Git-Internals-Git-Objects#Commit-Objects).

**Signature verification object**

The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:

These are the possible values for `reason` in the `verification` object:

| Value                    | Description                                                                                                                       |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| `expired_key`            | The key that made the signature is expired.                                                                                       |
| `not_signing_key`        | The "signing" flag is not among the usage flags in the GPG key that made the signature.                                           |
| `gpgverify_error`        | There was an error communicating with the signature verification service.                                                         |
| `gpgverify_unavailable`  | The signature verification service is currently unavailable.                                                                      |
| `unsigned`               | The object does not include a signature.                                                                                          |
| `unknown_signature_type` | A non-PGP signature was found in the commit.                                                                                      |
| `no_user`                | No user was associated with the `committer` email address in the commit.                                                          |
| `unverified_email`       | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. |
| `bad_email`              | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature.             |
| `unknown_key`            | The key that made the signature has not been registered with any user's account.                                                  |
| `malformed_signature`    | There was an error parsing the signature.                                                                                         |
| `invalid`                | The signature could not be cryptographically verified using the key whose key-id was found in the signature.                      |
| `valid`                  | None of the above errors applied, so the signature is considered to be verified.                                                  |

### parameters


| name | description |
|------|-------------|
| commit_sha | __Required__ commit_sha+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## git get-ref

https://developer.github.com/v3/git/refs/#get-a-reference

Returns a single reference from your Git database. The `:ref` in the URL must be formatted as `heads/<branch name>` for branches and `tags/<tag name>` for tags. If the `:ref` doesn't match an existing ref, a `404` is returned.

**Note:** You need to explicitly [request a pull request](https://developer.github.com/v3/pulls/#get-a-pull-request) to trigger a test merge commit, which checks the mergeability of pull requests. For more information, see "[Checking mergeability of pull requests](https://developer.github.com/v3/git/#checking-mergeability-of-pull-requests)".

### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## git get-tag

https://developer.github.com/v3/git/tags/#get-a-tag

**Signature verification object**

The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:

These are the possible values for `reason` in the `verification` object:

| Value                    | Description                                                                                                                       |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| `expired_key`            | The key that made the signature is expired.                                                                                       |
| `not_signing_key`        | The "signing" flag is not among the usage flags in the GPG key that made the signature.                                           |
| `gpgverify_error`        | There was an error communicating with the signature verification service.                                                         |
| `gpgverify_unavailable`  | The signature verification service is currently unavailable.                                                                      |
| `unsigned`               | The object does not include a signature.                                                                                          |
| `unknown_signature_type` | A non-PGP signature was found in the commit.                                                                                      |
| `no_user`                | No user was associated with the `committer` email address in the commit.                                                          |
| `unverified_email`       | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. |
| `bad_email`              | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature.             |
| `unknown_key`            | The key that made the signature has not been registered with any user's account.                                                  |
| `malformed_signature`    | There was an error parsing the signature.                                                                                         |
| `invalid`                | The signature could not be cryptographically verified using the key whose key-id was found in the signature.                      |
| `valid`                  | None of the above errors applied, so the signature is considered to be verified.                                                  |

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| tag_sha | __Required__ tag_sha parameter |

## git get-tree

https://developer.github.com/v3/git/trees/#get-a-tree

Returns a single tree using the SHA1 value for that tree.

If `truncated` is `true` in the response then the number of items in the `tree` array exceeded our maximum limit. If you need to fetch more items, use the non-recursive method of fetching trees, and fetch one sub-tree at a time.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| tree_sha | __Required__ tree_sha parameter |
| recursive | Setting this parameter to any value returns the objects or subtrees referenced by the tree specified in `:tree_sha`. For example, setting `recursive` to any of the following will enable returning objects or subtrees: `0`, `1`, `"true"`, and `"false"`. Omit this parameter to prevent recursively returning objects or subtrees. |

## git list-matching-refs

https://developer.github.com/v3/git/refs/#list-matching-references

Returns an array of references from your Git database that match the supplied name. The `:ref` in the URL must be formatted as `heads/<branch name>` for branches and `tags/<tag name>` for tags. If the `:ref` doesn't exist in the repository, but existing refs start with `:ref`, they will be returned as an array.

When you use this endpoint without providing a `:ref`, it will return an array of all the references from your Git database, including notes and stashes if they exist on the server. Anything in the namespace is returned, not just `heads` and `tags`.

**Note:** You need to explicitly [request a pull request](https://developer.github.com/v3/pulls/#get-a-pull-request) to trigger a test merge commit, which checks the mergeability of pull requests. For more information, see "[Checking mergeability of pull requests](https://developer.github.com/v3/git/#checking-mergeability-of-pull-requests)".

If you request matching references for a branch named `feature` but the branch `feature` doesn't exist, the response can still include other matching head refs that start with the word `feature`, such as `featureA` and `featureB`.

### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## git update-ref

https://developer.github.com/v3/git/refs/#update-a-reference



### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| sha | __Required__ The SHA1 value to set this reference to |
| force | Indicates whether to force the update or to make sure the update is a fast-forward update. Leaving this out or setting it to `false` will make sure you're not overwriting work. |

# gitignore


## gitignore get-all-templates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates

List all templates available to pass as an option when [creating a repository](https://developer.github.com/v3/repos/#create-a-repository-for-the-authenticated-user).

## gitignore get-template

https://developer.github.com/v3/gitignore/#get-a-gitignore-template

The API also allows fetching the source of a single template.
Use the raw [media type](https://developer.github.com/v3/media/) to get the raw contents.

### parameters


| name | description |
|------|-------------|
| name | __Required__ name parameter |

# interactions


## interactions get-restrictions-for-org

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization

Shows which group of GitHub users can interact with this organization and when the restriction expires. If there are no restrictions, you will see an empty response.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| sombra-preview | __Required__ The Interactions API is currently in public preview. See the [blog post](https://developer.github.com/changes/2018-12-18-interactions-preview) preview for more details. To access the API during the preview period, you must set this flag. |

## interactions get-restrictions-for-repo

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository

Shows which group of GitHub users can interact with this repository and when the restriction expires. If there are no restrictions, you will see an empty response.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| sombra-preview | __Required__ The Interactions API is currently in public preview. See the [blog post](https://developer.github.com/changes/2018-12-18-interactions-preview) preview for more details. To access the API during the preview period, you must set this flag. |

## interactions remove-restrictions-for-org

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization

Removes all interaction restrictions from public repositories in the given organization. You must be an organization owner to remove restrictions.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| sombra-preview | __Required__ The Interactions API is currently in public preview. See the [blog post](https://developer.github.com/changes/2018-12-18-interactions-preview) preview for more details. To access the API during the preview period, you must set this flag. |

## interactions remove-restrictions-for-repo

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository

Removes all interaction restrictions from the given repository. You must have owner or admin access to remove restrictions.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| sombra-preview | __Required__ The Interactions API is currently in public preview. See the [blog post](https://developer.github.com/changes/2018-12-18-interactions-preview) preview for more details. To access the API during the preview period, you must set this flag. |

## interactions set-restrictions-for-org

https://developer.github.com/v3/interactions/orgs/#set-interaction-restrictions-for-an-organization

Temporarily restricts interactions to certain GitHub users in any public repository in the given organization. You must be an organization owner to set these restrictions.

### parameters


| name | description |
|------|-------------|
| limit | __Required__ Specifies the group of GitHub users who can comment, open issues, or create pull requests in public repositories for the given organization. Must be one of: `existing_users`, `contributors_only`, or `collaborators_only`. |
| org | __Required__  |
| sombra-preview | __Required__ The Interactions API is currently in public preview. See the [blog post](https://developer.github.com/changes/2018-12-18-interactions-preview) preview for more details. To access the API during the preview period, you must set this flag. |

## interactions set-restrictions-for-repo

https://developer.github.com/v3/interactions/repos/#set-interaction-restrictions-for-a-repository

Temporarily restricts interactions to certain GitHub users within the given repository. You must have owner or admin access to set restrictions.

### parameters


| name | description |
|------|-------------|
| limit | __Required__ Specifies the group of GitHub users who can comment, open issues, or create pull requests for the given repository. Must be one of: `existing_users`, `contributors_only`, or `collaborators_only`. |
| repo | __Required__ repository in OWNER/REPO form |
| sombra-preview | __Required__ The Interactions API is currently in public preview. See the [blog post](https://developer.github.com/changes/2018-12-18-interactions-preview) preview for more details. To access the API during the preview period, you must set this flag. |

# issues


## issues add-assignees

https://developer.github.com/v3/issues/assignees/#add-assignees-to-an-issue

Adds up to 10 assignees to an issue. Users already assigned to an issue are not replaced.

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| assignees | Usernames of people to assign this issue to. _NOTE: Only users with push access can add assignees to an issue. Assignees are silently ignored otherwise._ |

## issues add-labels

https://developer.github.com/v3/issues/labels/#add-labels-to-an-issue



### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| labels | __Required__ The name of the label to add to the issue. Must contain at least one label. **Note:** Alternatively, you can pass a single label as a `string` or an `array` of labels directly, but GitHub recommends passing an object with the `labels` key. |
| repo | __Required__ repository in OWNER/REPO form |

## issues check-user-can-be-assigned

https://developer.github.com/v3/issues/assignees/#check-if-a-user-can-be-assigned

Checks if a user has permission to be assigned to an issue in this repository.

If the `assignee` can be assigned to issues in the repository, a `204` header with no content is returned.

Otherwise a `404` status code is returned.

### parameters


| name | description |
|------|-------------|
| assignee | __Required__ assignee parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues create

https://developer.github.com/v3/issues/#create-an-issue

Any user with pull access to a repository can create an issue. If [issues are disabled in the repository](https://help.github.com/articles/disabling-issues/), the API returns a `410 Gone` status.

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| title | __Required__ The title of the issue. |
| assignee | Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_ |
| assignees | Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._ |
| body | The contents of the issue. |
| labels | Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._ |
| milestone | The `number` of the milestone to associate this issue with. _NOTE: Only users with push access can set the milestone for new issues. The milestone is silently dropped otherwise._ |

## issues create-comment

https://developer.github.com/v3/issues/comments/#create-an-issue-comment

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The contents of the comment. |
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues create-label

https://developer.github.com/v3/issues/labels/#create-a-label



### parameters


| name | description |
|------|-------------|
| name | __Required__ The name of the label. Emoji can be added to label names, using either native emoji or colon-style markup. For example, typing `:strawberry:` will render the emoji ![:strawberry:](https://github.githubassets.com/images/icons/emoji/unicode/1f353.png ":strawberry:"). For a full list of available emoji and codes, see [emoji-cheat-sheet.com](http://emoji-cheat-sheet.com/). |
| repo | __Required__ repository in OWNER/REPO form |
| color | The [hexadecimal color code](http://www.color-hex.com/) for the label, without the leading `#`. |
| description | A short description of the label. |

## issues create-milestone

https://developer.github.com/v3/issues/milestones/#create-a-milestone



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| title | __Required__ The title of the milestone. |
| description | A description of the milestone. |
| due_on | The milestone due date. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| state | The state of the milestone. Either `open` or `closed`. |

## issues delete-comment

https://developer.github.com/v3/issues/comments/#delete-an-issue-comment



### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues delete-label

https://developer.github.com/v3/issues/labels/#delete-a-label



### parameters


| name | description |
|------|-------------|
| name | __Required__ name parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues delete-milestone

https://developer.github.com/v3/issues/milestones/#delete-a-milestone



### parameters


| name | description |
|------|-------------|
| milestone_number | __Required__ milestone_number parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues get

https://developer.github.com/v3/issues/#get-an-issue

The API returns a [`301 Moved Permanently` status](https://developer.github.com/v3/#http-redirects) if the issue was
[transferred](https://help.github.com/articles/transferring-an-issue-to-another-repository/) to another repository. If
the issue was transferred to or deleted from a repository where the authenticated user lacks read access, the API
returns a `404 Not Found` status. If the issue was deleted from a repository where the authenticated user has read
access, the API returns a `410 Gone` status. To receive webhook events for transferred and deleted issues, subscribe
to the [`issues`](https://developer.github.com/webhooks/event-payloads/#issues) webhook.

**Note**: GitHub's REST API v3 considers every pull request an issue, but not every issue is a pull request. For this
reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by
the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull
request id, use the "[List pull requests](https://developer.github.com/v3/pulls/#list-pull-requests)" endpoint.

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## issues get-comment

https://developer.github.com/v3/issues/comments/#get-an-issue-comment



### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| machine-man-preview | To access the API with your GitHub App, you must set this flag for your requests. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## issues get-event

https://developer.github.com/v3/issues/events/#get-an-issue-event



### parameters


| name | description |
|------|-------------|
| event_id | __Required__ event_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| machine-man-preview | To access the API with your GitHub App, you must set this flag for your requests. |
| sailor-v-preview | You can now use the REST API to add a reason when you lock an issue, and you will see lock reasons in responses that include issues or pull requests. You will also see lock reasons in `locked` events. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-01-10-lock-reason-api-preview) for full details. To access this feature, you must set this flag. |
| starfox-preview | Project card details are now shown in REST API v3 responses for project-related issue and timeline events. This feature is now available for developers to preview. For details, see the [blog post](https://developer.github.com/changes/2018-09-05-project-card-events).<br><br>To receive the `project_card` attribute, project boards must be [enabled](https://help.github.com/articles/disabling-project-boards-in-a-repository) for a repository, and you must set this flag. |

## issues get-label

https://developer.github.com/v3/issues/labels/#get-a-label



### parameters


| name | description |
|------|-------------|
| name | __Required__ name parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues get-milestone

https://developer.github.com/v3/issues/milestones/#get-a-milestone



### parameters


| name | description |
|------|-------------|
| milestone_number | __Required__ milestone_number parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues list

https://developer.github.com/v3/issues/#list-issues-assigned-to-the-authenticated-user

List issues assigned to the authenticated user across all visible repositories including owned repositories, member
repositories, and organization repositories. You can use the `filter` query parameter to fetch issues that are not
necessarily assigned to you. See the [Parameters table](https://developer.github.com/v3/issues/#parameters) for more
information.


**Note**: GitHub's REST API v3 considers every pull request an issue, but not every issue is a pull request. For this
reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by
the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull
request id, use the "[List pull requests](https://developer.github.com/v3/pulls/#list-pull-requests)" endpoint.

### parameters


| name | description |
|------|-------------|
| collab |  |
| direction | One of `asc` (ascending) or `desc` (descending). |
| filter | Indicates which sorts of issues to return. Can be one of:  <br>\* `assigned`: Issues assigned to you  <br>\* `created`: Issues created by you  <br>\* `mentioned`: Issues mentioning you  <br>\* `subscribed`: Issues you're subscribed to updates for  <br>\* `all`: All issues the authenticated user can see, regardless of participation or creation |
| labels | A list of comma separated label names. Example: `bug,ui,@high` |
| machine-man-preview | To access the API with your GitHub App, you must set this flag for your requests. |
| orgs |  |
| owned |  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| pulls |  |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| sort | What to sort results by. Can be either `created`, `updated`, `comments`. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| state | Indicates the state of the issues to return. Can be either `open`, `closed`, or `all`. |

## issues list-assignees

https://developer.github.com/v3/issues/assignees/#list-assignees

Lists the [available assignees](https://help.github.com/articles/assigning-issues-and-pull-requests-to-other-github-users/) for issues in a repository.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## issues list-comments

https://developer.github.com/v3/issues/comments/#list-issue-comments

Issue Comments are ordered by ascending ID.

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## issues list-comments-for-repo

https://developer.github.com/v3/issues/comments/#list-issue-comments-for-a-repository

By default, Issue Comments are ordered by ascending ID.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| direction | Either `asc` or `desc`. Ignored without the `sort` parameter. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| sort | One of `created` (when the repository was starred) or `updated` (when it was last pushed to). |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## issues list-events

https://developer.github.com/v3/issues/events/#list-issue-events



### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sailor-v-preview | You can now use the REST API to add a reason when you lock an issue, and you will see lock reasons in responses that include issues or pull requests. You will also see lock reasons in `locked` events. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-01-10-lock-reason-api-preview) for full details. To access this feature, you must set this flag. |
| starfox-preview | Project card details are now shown in REST API v3 responses for project-related issue and timeline events. This feature is now available for developers to preview. For details, see the [blog post](https://developer.github.com/changes/2018-09-05-project-card-events).<br><br>To receive the `project_card` attribute, project boards must be [enabled](https://help.github.com/articles/disabling-project-boards-in-a-repository) for a repository, and you must set this flag. |

## issues list-events-for-repo

https://developer.github.com/v3/issues/events/#list-issue-events-for-a-repository



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sailor-v-preview | You can now use the REST API to add a reason when you lock an issue, and you will see lock reasons in responses that include issues or pull requests. You will also see lock reasons in `locked` events. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-01-10-lock-reason-api-preview) for full details. To access this feature, you must set this flag. |
| starfox-preview | Project card details are now shown in REST API v3 responses for project-related issue and timeline events. This feature is now available for developers to preview. For details, see the [blog post](https://developer.github.com/changes/2018-09-05-project-card-events).<br><br>To receive the `project_card` attribute, project boards must be [enabled](https://help.github.com/articles/disabling-project-boards-in-a-repository) for a repository, and you must set this flag. |

## issues list-events-for-timeline

https://developer.github.com/v3/issues/timeline/#list-timeline-events-for-an-issue



### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| mockingbird-preview | __Required__ The API to get issue timeline events is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-23-timeline-preview-api/) for full details. To access the API you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| starfox-preview | Project card details are now shown in REST API v3 responses for project-related issue and timeline events. This feature is now available for developers to preview. For details, see the [blog post](https://developer.github.com/changes/2018-09-05-project-card-events).<br><br>To receive the `project_card` attribute, project boards must be [enabled](https://help.github.com/articles/disabling-project-boards-in-a-repository) for a repository, and you must set this flag. |

## issues list-for-authenticated-user

https://developer.github.com/v3/issues/#list-user-account-issues-assigned-to-the-authenticated-user

List issues across owned and member repositories assigned to the authenticated user.

**Note**: GitHub's REST API v3 considers every pull request an issue, but not every issue is a pull request. For this
reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by
the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull
request id, use the "[List pull requests](https://developer.github.com/v3/pulls/#list-pull-requests)" endpoint.

### parameters


| name | description |
|------|-------------|
| direction | One of `asc` (ascending) or `desc` (descending). |
| filter | Indicates which sorts of issues to return. Can be one of:  <br>\* `assigned`: Issues assigned to you  <br>\* `created`: Issues created by you  <br>\* `mentioned`: Issues mentioning you  <br>\* `subscribed`: Issues you're subscribed to updates for  <br>\* `all`: All issues the authenticated user can see, regardless of participation or creation |
| labels | A list of comma separated label names. Example: `bug,ui,@high` |
| machine-man-preview | To access the API with your GitHub App, you must set this flag for your requests. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| sort | What to sort results by. Can be either `created`, `updated`, `comments`. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| state | Indicates the state of the issues to return. Can be either `open`, `closed`, or `all`. |

## issues list-for-org

https://developer.github.com/v3/issues/#list-organization-issues-assigned-to-the-authenticated-user

List issues in an organization assigned to the authenticated user.

**Note**: GitHub's REST API v3 considers every pull request an issue, but not every issue is a pull request. For this
reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by
the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull
request id, use the "[List pull requests](https://developer.github.com/v3/pulls/#list-pull-requests)" endpoint.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| direction | One of `asc` (ascending) or `desc` (descending). |
| filter | Indicates which sorts of issues to return. Can be one of:  <br>\* `assigned`: Issues assigned to you  <br>\* `created`: Issues created by you  <br>\* `mentioned`: Issues mentioning you  <br>\* `subscribed`: Issues you're subscribed to updates for  <br>\* `all`: All issues the authenticated user can see, regardless of participation or creation |
| labels | A list of comma separated label names. Example: `bug,ui,@high` |
| machine-man-preview | To access the API with your GitHub App, you must set this flag for your requests. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| sort | What to sort results by. Can be either `created`, `updated`, `comments`. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| state | Indicates the state of the issues to return. Can be either `open`, `closed`, or `all`. |

## issues list-for-repo

https://developer.github.com/v3/issues/#list-repository-issues

List issues in a repository.

**Note**: GitHub's REST API v3 considers every pull request an issue, but not every issue is a pull request. For this
reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by
the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull
request id, use the "[List pull requests](https://developer.github.com/v3/pulls/#list-pull-requests)" endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| assignee | Can be the name of a user. Pass in `none` for issues with no assigned user, and `*` for issues assigned to any user. |
| creator | The user that created the issue. |
| direction | One of `asc` (ascending) or `desc` (descending). |
| labels | A list of comma separated label names. Example: `bug,ui,@high` |
| machine-man-preview | To access the API with your GitHub App, you must set this flag for your requests. |
| mentioned | A user that's mentioned in the issue. |
| milestone | If an `integer` is passed, it should refer to a milestone by its `number` field. If the string `*` is passed, issues with any milestone are accepted. If the string `none` is passed, issues without milestones are returned. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| sort | What to sort results by. Can be either `created`, `updated`, `comments`. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| state | Indicates the state of the issues to return. Can be either `open`, `closed`, or `all`. |

## issues list-labels-for-milestone

https://developer.github.com/v3/issues/labels/#list-labels-for-issues-in-a-milestone



### parameters


| name | description |
|------|-------------|
| milestone_number | __Required__ milestone_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## issues list-labels-for-repo

https://developer.github.com/v3/issues/labels/#list-labels-for-a-repository



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## issues list-labels-on-issue

https://developer.github.com/v3/issues/labels/#list-labels-for-an-issue



### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## issues list-milestones

https://developer.github.com/v3/issues/milestones/#list-milestones



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| direction | The direction of the sort. Either `asc` or `desc`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | What to sort results by. Either `due_on` or `completeness`. |
| state | The state of the milestone. Either `open`, `closed`, or `all`. |

## issues lock

https://developer.github.com/v3/issues/#lock-an-issue

Users with push access can lock an issue or pull request's conversation.

Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| lock_reason | The reason for locking the issue or pull request conversation. Lock will fail if you don't use one of these reasons:  <br>\* `off-topic`  <br>\* `too heated`  <br>\* `resolved`  <br>\* `spam` |
| sailor-v-preview | You can now use the REST API to add a reason when you lock an issue, and you will see lock reasons in responses that include issues or pull requests. You will also see lock reasons in `locked` events. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-01-10-lock-reason-api-preview) for full details. To access this feature, you must set this flag. |

## issues remove-all-labels

https://developer.github.com/v3/issues/labels/#remove-all-labels-from-an-issue



### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues remove-assignees

https://developer.github.com/v3/issues/assignees/#remove-assignees-from-an-issue

Removes one or more assignees from an issue.

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| assignees | Usernames of assignees to remove from an issue. _NOTE: Only users with push access can remove assignees from an issue. Assignees are silently ignored otherwise._ |

## issues remove-label

https://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue

Removes the specified label from the issue, and returns the remaining labels on the issue. This endpoint returns a `404 Not Found` status if the label does not exist.

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| name | __Required__ name parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues set-labels

https://developer.github.com/v3/issues/labels/#set-labels-for-an-issue

Removes any previous labels and sets the new labels for an issue.

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| labels | The names of the labels to add to the issue. You can pass an empty array to remove all labels. **Note:** Alternatively, you can pass a single label as a `string` or an `array` of labels directly, but GitHub recommends passing an object with the `labels` key. |

## issues unlock

https://developer.github.com/v3/issues/#unlock-an-issue

Users with push access can unlock an issue's conversation.

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues update

https://developer.github.com/v3/issues/#update-an-issue

Issue owners and users with push access can edit an issue.

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| assignee | Login for the user that this issue should be assigned to. **This field is deprecated.** |
| assignees | Logins for Users to assign to this issue. Pass one or more user logins to _replace_ the set of assignees on this Issue. Send an empty array (`[]`) to clear all assignees from the Issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._ |
| body | The contents of the issue. |
| labels | Labels to associate with this issue. Pass one or more Labels to _replace_ the set of Labels on this Issue. Send an empty array (`[]`) to clear all Labels from the Issue. _NOTE: Only users with push access can set labels for issues. Labels are silently dropped otherwise._ |
| milestone | The `number` of the milestone to associate this issue with or `null` to remove current. _NOTE: Only users with push access can set the milestone for issues. The milestone is silently dropped otherwise._ |
| state | State of the issue. Either `open` or `closed`. |
| title | The title of the issue. |

## issues update-comment

https://developer.github.com/v3/issues/comments/#update-an-issue-comment



### parameters


| name | description |
|------|-------------|
| body | __Required__ The contents of the comment. |
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## issues update-label

https://developer.github.com/v3/issues/labels/#update-a-label



### parameters


| name | description |
|------|-------------|
| name | __Required__ name parameter |
| repo | __Required__ repository in OWNER/REPO form |
| color | The [hexadecimal color code](http://www.color-hex.com/) for the label, without the leading `#`. |
| description | A short description of the label. |
| new_name | The new name of the label. Emoji can be added to label names, using either native emoji or colon-style markup. For example, typing `:strawberry:` will render the emoji ![:strawberry:](https://github.githubassets.com/images/icons/emoji/unicode/1f353.png ":strawberry:"). For a full list of available emoji and codes, see [emoji-cheat-sheet.com](http://emoji-cheat-sheet.com/). |

## issues update-milestone

https://developer.github.com/v3/issues/milestones/#update-a-milestone



### parameters


| name | description |
|------|-------------|
| milestone_number | __Required__ milestone_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| description | A description of the milestone. |
| due_on | The milestone due date. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| state | The state of the milestone. Either `open` or `closed`. |
| title | The title of the milestone. |

# licenses


## licenses get

https://developer.github.com/v3/licenses/#get-a-license



### parameters


| name | description |
|------|-------------|
| license | __Required__ license parameter |

## licenses get-all-commonly-used

https://developer.github.com/v3/licenses/#get-all-commonly-used-licenses



### parameters


| name | description |
|------|-------------|
| featured |  |
| per_page | Results per page (max 100) |

## licenses get-for-repo

https://developer.github.com/v3/licenses/#get-the-license-for-a-repository

This method returns the contents of the repository's license file, if one is detected.

Similar to [Get repository content](https://developer.github.com/v3/repos/contents/#get-repository-content), this method also supports [custom media types](https://developer.github.com/v3/repos/contents/#custom-media-types) for retrieving the raw license content or rendered license HTML.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

# markdown


## markdown render

https://developer.github.com/v3/markdown/#render-a-markdown-document



### parameters


| name | description |
|------|-------------|
| text | __Required__ The Markdown text to render in HTML. |
| context | The repository context to use when creating references in `gfm` mode. |
| mode | The rendering mode. |

## markdown render-raw

https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode

You must send Markdown as plain text (using a `Content-Type` header of `text/plain` or `text/x-markdown`) to this endpoint, rather than using JSON format. In raw mode, [GitHub Flavored Markdown](https://github.github.com/gfm/) is not supported and Markdown will be rendered in plain format like a README.md file. Markdown content must be 400 KB or less.

### parameters


| name | description |
|------|-------------|
| file | __Required__ the file to upload |

# meta


## meta get

https://developer.github.com/v3/meta/#get-github-meta-information

This endpoint provides a list of GitHub's IP addresses. For more information, see "[About GitHub's IP addresses](https://help.github.com/articles/about-github-s-ip-addresses/)."

## meta get-octocat



Get the octocat as ASCII art

### parameters


| name | description |
|------|-------------|
| s | The words to show in Octocat's speech bubble |

## meta get-zen



Get a random sentence from the Zen of GitHub

## meta root



Get Hypermedia links to resources accessible in GitHub's REST API

# migrations


## migrations cancel-import

https://developer.github.com/v3/migrations/source_imports/#cancel-an-import

Stop an import for a repository.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## migrations delete-archive-for-authenticated-user

https://developer.github.com/v3/migrations/users/#delete-a-user-migration-archive

Deletes a previous migration archive. Downloadable migration archives are automatically deleted after seven days. Migration metadata, which is returned in the [List user migrations](https://developer.github.com/v3/migrations/users/#list-user-migrations) and [Get a user migration status](https://developer.github.com/v3/migrations/users/#get-a-user-migration-status) endpoints, will continue to be available even after an archive is deleted.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |

## migrations delete-archive-for-org

https://developer.github.com/v3/migrations/orgs/#delete-an-organization-migration-archive

Deletes a previous migration archive. Migration archives are automatically deleted after seven days.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| org | __Required__  |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |

## migrations download-archive-for-org

https://developer.github.com/v3/migrations/orgs/#download-an-organization-migration-archive

Fetches the URL to a migration archive.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| org | __Required__  |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |

## migrations get-archive-for-authenticated-user

https://developer.github.com/v3/migrations/users/#download-a-user-migration-archive

Fetches the URL to download the migration archive as a `tar.gz` file. Depending on the resources your repository uses, the migration archive can contain JSON files with data for these objects:

*   attachments
*   bases
*   commit\_comments
*   issue\_comments
*   issue\_events
*   issues
*   milestones
*   organizations
*   projects
*   protected\_branches
*   pull\_request\_reviews
*   pull\_requests
*   releases
*   repositories
*   review\_comments
*   schema
*   users

The archive will also contain an `attachments` directory that includes all attachment files uploaded to GitHub.com and a `repositories` directory that contains the repository's Git data.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |

## migrations get-commit-authors

https://developer.github.com/v3/migrations/source_imports/#get-commit-authors

Each type of source control system represents authors in a different way. For example, a Git commit author has a display name and an email address, but a Subversion commit author just has a username. The GitHub Importer will make the author information valid, but the author might not be correct. For example, it will change the bare Subversion username `hubot` into something like `hubot <hubot@12341234-abab-fefe-8787-fedcba987654>`.

This endpoint and the [Map a commit author](https://developer.github.com/v3/migrations/source_imports/#map-a-commit-author) endpoint allow you to provide correct Git author information.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## migrations get-import-status

https://developer.github.com/v3/migrations/source_imports/#get-an-import-status

View the progress of an import.

**Import status**

This section includes details about the possible values of the `status` field of the Import Progress response.

An import that does not have errors will progress through these steps:

*   `detecting` - the "detection" step of the import is in progress because the request did not include a `vcs` parameter. The import is identifying the type of source control present at the URL.
*   `importing` - the "raw" step of the import is in progress. This is where commit data is fetched from the original repository. The import progress response will include `commit_count` (the total number of raw commits that will be imported) and `percent` (0 - 100, the current progress through the import).
*   `mapping` - the "rewrite" step of the import is in progress. This is where SVN branches are converted to Git branches, and where author updates are applied. The import progress response does not include progress information.
*   `pushing` - the "push" step of the import is in progress. This is where the importer updates the repository on GitHub. The import progress response will include `push_percent`, which is the percent value reported by `git push` when it is "Writing objects".
*   `complete` - the import is complete, and the repository is ready on GitHub.

If there are problems, you will see one of these in the `status` field:

*   `auth_failed` - the import requires authentication in order to connect to the original repository. To update authentication for the import, please see the [Update an import](https://developer.github.com/v3/migrations/source_imports/#update-an-import) section.
*   `error` - the import encountered an error. The import progress response will include the `failed_step` and an error message. Contact [GitHub Support](https://github.com/contact) or [GitHub Premium Support](https://premium.githubsupport.com) for more information.
*   `detection_needs_auth` - the importer requires authentication for the originating repository to continue detection. To update authentication for the import, please see the [Update an import](https://developer.github.com/v3/migrations/source_imports/#update-an-import) section.
*   `detection_found_nothing` - the importer didn't recognize any source control at the URL. To resolve, [Cancel the import](https://developer.github.com/v3/migrations/source_imports/#cancel-an-import) and [retry](https://developer.github.com/v3/migrations/source_imports/#start-an-import) with the correct URL.
*   `detection_found_multiple` - the importer found several projects or repositories at the provided URL. When this is the case, the Import Progress response will also include a `project_choices` field with the possible project choices as values. To update project choice, please see the [Update an import](https://developer.github.com/v3/migrations/source_imports/#update-an-import) section.

**The project_choices field**

When multiple projects are found at the provided URL, the response hash will include a `project_choices` field, the value of which is an array of hashes each representing a project choice. The exact key/value pairs of the project hashes will differ depending on the version control type.

**Git LFS related fields**

This section includes details about Git LFS related fields that may be present in the Import Progress response.

*   `use_lfs` - describes whether the import has been opted in or out of using Git LFS. The value can be `opt_in`, `opt_out`, or `undecided` if no action has been taken.
*   `has_large_files` - the boolean value describing whether files larger than 100MB were found during the `importing` step.
*   `large_files_size` - the total size in gigabytes of files larger than 100MB found in the originating repository.
*   `large_files_count` - the total number of files larger than 100MB found in the originating repository. To see a list of these files, make a "Get Large Files" request.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## migrations get-large-files

https://developer.github.com/v3/migrations/source_imports/#get-large-files

List files larger than 100MB found during the import

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## migrations get-status-for-authenticated-user

https://developer.github.com/v3/migrations/users/#get-a-user-migration-status

Fetches a single user migration. The response includes the `state` of the migration, which can be one of the following values:

*   `pending` - the migration hasn't started yet.
*   `exporting` - the migration is in progress.
*   `exported` - the migration finished successfully.
*   `failed` - the migration failed.

Once the migration has been `exported` you can [download the migration archive](https://developer.github.com/v3/migrations/users/#download-a-user-migration-archive).

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |
| exclude |  |

## migrations get-status-for-org

https://developer.github.com/v3/migrations/orgs/#get-an-organization-migration-status

Fetches the status of a migration.

The `state` of a migration can be one of the following values:

*   `pending`, which means the migration hasn't started yet.
*   `exporting`, which means the migration is in progress.
*   `exported`, which means the migration finished successfully.
*   `failed`, which means the migration failed.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| org | __Required__  |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |

## migrations list-for-authenticated-user

https://developer.github.com/v3/migrations/users/#list-user-migrations

Lists all migrations a user has started.

### parameters


| name | description |
|------|-------------|
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## migrations list-for-org

https://developer.github.com/v3/migrations/orgs/#list-organization-migrations

Lists the most recent migrations.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## migrations list-repos-for-org

https://developer.github.com/v3/migrations/orgs/#list-repositories-in-an-organization-migration

List all the repositories for this organization migration.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| org | __Required__  |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## migrations list-repos-for-user

https://developer.github.com/v3/migrations/users/#list-repositories-for-a-user-migration

Lists all the repositories for this user migration.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## migrations map-commit-author

https://developer.github.com/v3/migrations/source_imports/#map-a-commit-author

Update an author's identity for the import. Your application can continue updating authors any time before you push new commits to the repository.

### parameters


| name | description |
|------|-------------|
| author_id | __Required__ author_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| email | The new Git author email. |
| name | The new Git author name. |
| remote_id |  |

## migrations set-lfs-preference

https://developer.github.com/v3/migrations/source_imports/#update-git-lfs-preference

You can import repositories from Subversion, Mercurial, and TFS that include files larger than 100MB. This ability is powered by [Git LFS](https://git-lfs.github.com). You can learn more about our LFS feature and working with large files [on our help site](https://help.github.com/articles/versioning-large-files/).

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| use_lfs | __Required__ Can be one of `opt_in` (large files will be stored using Git LFS) or `opt_out` (large files will be removed during the import). |

## migrations start-for-authenticated-user

https://developer.github.com/v3/migrations/users/#start-a-user-migration

Initiates the generation of a user migration archive.

### parameters


| name | description |
|------|-------------|
| repositories | __Required__  |
| exclude | Exclude attributes from the API response to improve performance |
| exclude_attachments | Do not include attachments in the migration |
| lock_repositories | Lock the repositories being migrated at the start of the migration |

## migrations start-for-org

https://developer.github.com/v3/migrations/orgs/#start-an-organization-migration

Initiates the generation of a migration archive.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| repositories | __Required__ A list of arrays indicating which repositories should be migrated. |
| exclude |  |
| exclude_attachments | Indicates whether attachments should be excluded from the migration (to reduce migration archive file size). |
| lock_repositories | Indicates whether repositories should be locked (to prevent manipulation) while migrating data. |

## migrations start-import

https://developer.github.com/v3/migrations/source_imports/#start-an-import

Start a source import to a GitHub repository using GitHub Importer.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| vcs_url | __Required__ The URL of the originating repository. |
| tfvc_project | For a tfvc import, the name of the project that is being imported. |
| vcs | The originating VCS type. Can be one of `subversion`, `git`, `mercurial`, or `tfvc`. Please be aware that without this parameter, the import job will take additional time to detect the VCS type before beginning the import. This detection step will be reflected in the response. |
| vcs_password | If authentication is required, the password to provide to `vcs_url`. |
| vcs_username | If authentication is required, the username to provide to `vcs_url`. |

## migrations unlock-repo-for-authenticated-user

https://developer.github.com/v3/migrations/users/#unlock-a-user-repository

Unlocks a repository. You can lock repositories when you [start a user migration](https://developer.github.com/v3/migrations/users/#start-a-user-migration). Once the migration is complete you can unlock each repository to begin using it again or [delete the repository](https://developer.github.com/v3/repos/#delete-a-repository) if you no longer need the source data. Returns a status of `404 Not Found` if the repository is not locked.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| repo_name | __Required__ repo_name parameter |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |

## migrations unlock-repo-for-org

https://developer.github.com/v3/migrations/orgs/#unlock-an-organization-repository

Unlocks a repository that was locked for migration. You should unlock each migrated repository and [delete them](https://developer.github.com/v3/repos/#delete-a-repository) when the migration is complete and you no longer need the source data.

### parameters


| name | description |
|------|-------------|
| migration_id | __Required__ migration_id parameter |
| org | __Required__  |
| repo_name | __Required__ repo_name parameter |
| wyandotte-preview | __Required__ To access the Migrations API, you must set this flag. |

## migrations update-import

https://developer.github.com/v3/migrations/source_imports/#update-an-import

An import can be updated with credentials or a project choice by passing in the appropriate parameters in this API
request. If no parameters are provided, the import will be restarted.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| tfvc_project |  |
| vcs |  |
| vcs_password | The password to provide to the originating repository. |
| vcs_username | The username to provide to the originating repository. |

# oauth-authorizations


## oauth-authorizations create-authorization

https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

**Warning:** Apps must use the [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow) to obtain OAuth tokens that work with GitHub SAML organizations. OAuth tokens created using the Authorizations API will be unable to access GitHub SAML organizations. For more information, see the [blog post](https://developer.github.com/changes/2019-11-05-deprecated-passwords-and-authorizations-api).

Creates OAuth tokens using [Basic Authentication](https://developer.github.com/v3/auth#basic-authentication). If you have two-factor authentication setup, Basic Authentication for this endpoint requires that you use a one-time password (OTP) and your username and password instead of tokens. For more information, see "[Working with two-factor authentication](https://developer.github.com/v3/auth/#working-with-two-factor-authentication)."

To create tokens for a particular OAuth application using this endpoint, you must authenticate as the user you want to create an authorization for and provide the app's client ID and secret, found on your OAuth application's settings page. If your OAuth application intends to create multiple tokens for one user, use `fingerprint` to differentiate between them.

You can also create tokens on GitHub from the [personal access tokens settings](https://github.com/settings/tokens) page. Read more about these tokens in [the GitHub Help documentation](https://help.github.com/articles/creating-an-access-token-for-command-line-use).

Organizations that enforce SAML SSO require personal access tokens to be allowed. Read more about allowing tokens in [the GitHub Help documentation](https://help.github.com/articles/about-identity-and-access-management-with-saml-single-sign-on).

### parameters


| name | description |
|------|-------------|
| client_id | The OAuth app client key for which to create the token. |
| client_secret | The OAuth app client secret for which to create the token. |
| fingerprint | A unique string to distinguish an authorization from others created for the same client ID and user. |
| note | A note to remind you what the OAuth token is for. |
| note_url | A URL to remind you what app the OAuth token is for. |
| scopes | A list of scopes that this authorization is in. |

## oauth-authorizations delete-authorization

https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

### parameters


| name | description |
|------|-------------|
| authorization_id | __Required__ authorization_id parameter |

## oauth-authorizations delete-grant

https://developer.github.com/v3/oauth_authorizations/#delete-a-grant

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

Deleting an OAuth application's grant will also delete all OAuth tokens associated with the application for your user. Once deleted, the application has no access to your account and is no longer listed on [the application authorizations settings screen within GitHub](https://github.com/settings/applications#authorized).

### parameters


| name | description |
|------|-------------|
| grant_id | __Required__ grant_id parameter |

## oauth-authorizations get-authorization

https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

### parameters


| name | description |
|------|-------------|
| authorization_id | __Required__ authorization_id parameter |

## oauth-authorizations get-grant

https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

### parameters


| name | description |
|------|-------------|
| grant_id | __Required__ grant_id parameter |

## oauth-authorizations get-or-create-authorization-for-app

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

**Warning:** Apps must use the [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow) to obtain OAuth tokens that work with GitHub SAML organizations. OAuth tokens created using the Authorizations API will be unable to access GitHub SAML organizations. For more information, see the [blog post](https://developer.github.com/changes/2019-11-05-deprecated-passwords-and-authorizations-api).

Creates a new authorization for the specified OAuth application, only if an authorization for that application doesn't already exist for the user. The URL includes the 20 character client ID for the OAuth app that is requesting the token. It returns the user's existing authorization for the application if one is present. Otherwise, it creates and returns a new one.

If you have two-factor authentication setup, Basic Authentication for this endpoint requires that you use a one-time password (OTP) and your username and password instead of tokens. For more information, see "[Working with two-factor authentication](https://developer.github.com/v3/auth/#working-with-two-factor-authentication)."

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

### parameters


| name | description |
|------|-------------|
| client_id | __Required__  |
| client_secret | __Required__ The OAuth app client secret for which to create the token. |
| fingerprint | A unique string to distinguish an authorization from others created for the same client ID and user. |
| note | A note to remind you what the OAuth token is for. |
| note_url | A URL to remind you what app the OAuth token is for. |
| scopes | A list of scopes that this authorization is in. |

## oauth-authorizations get-or-create-authorization-for-app-and-fingerprint

https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

**Warning:** Apps must use the [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow) to obtain OAuth tokens that work with GitHub SAML organizations. OAuth tokens created using the Authorizations API will be unable to access GitHub SAML organizations. For more information, see the [blog post](https://developer.github.com/changes/2019-11-05-deprecated-passwords-and-authorizations-api).

This method will create a new authorization for the specified OAuth application, only if an authorization for that application and fingerprint do not already exist for the user. The URL includes the 20 character client ID for the OAuth app that is requesting the token. `fingerprint` is a unique string to distinguish an authorization from others created for the same client ID and user. It returns the user's existing authorization for the application if one is present. Otherwise, it creates and returns a new one.

If you have two-factor authentication setup, Basic Authentication for this endpoint requires that you use a one-time password (OTP) and your username and password instead of tokens. For more information, see "[Working with two-factor authentication](https://developer.github.com/v3/auth/#working-with-two-factor-authentication)."

### parameters


| name | description |
|------|-------------|
| client_id | __Required__  |
| client_secret | __Required__ The OAuth app client secret for which to create the token. |
| fingerprint | __Required__ fingerprint parameter |
| note | A note to remind you what the OAuth token is for. |
| note_url | A URL to remind you what app the OAuth token is for. |
| scopes | A list of scopes that this authorization is in. |

## oauth-authorizations list-authorizations

https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## oauth-authorizations list-grants

https://developer.github.com/v3/oauth_authorizations/#list-your-grants

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

You can use this API to list the set of OAuth applications that have been granted access to your account. Unlike the [list your authorizations](https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations) API, this API does not manage individual tokens. This API will return one entry for each OAuth application that has been granted access to your account, regardless of the number of tokens an application has generated for your user. The list of OAuth applications returned matches what is shown on [the application authorizations settings screen within GitHub](https://github.com/settings/applications#authorized). The `scopes` returned are the union of scopes authorized for the application. For example, if an application has one token with `repo` scope and another token with `user` scope, the grant will return `["repo", "user"]`.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## oauth-authorizations update-authorization

https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization

**Deprecation Notice:** GitHub will discontinue the [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/), which is used by integrations to create personal access tokens and OAuth tokens, and you must now create these tokens using our [web application flow](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow). The [OAuth Authorizations API](https://developer.github.com/v3/oauth_authorizations/) will be removed on November, 13, 2020. For more information, including scheduled brownouts, see the [blog post](https://developer.github.com/changes/2020-02-14-deprecating-oauth-auth-endpoint/).

If you have two-factor authentication setup, Basic Authentication for this endpoint requires that you use a one-time password (OTP) and your username and password instead of tokens. For more information, see "[Working with two-factor authentication](https://developer.github.com/v3/auth/#working-with-two-factor-authentication)."

You can only send one of these scope keys at a time.

### parameters


| name | description |
|------|-------------|
| authorization_id | __Required__ authorization_id parameter |
| add_scopes | A list of scopes to add to this authorization. |
| fingerprint | A unique string to distinguish an authorization from others created for the same client ID and user. |
| note | A note to remind you what the OAuth token is for. |
| note_url | A URL to remind you what app the OAuth token is for. |
| remove_scopes | A list of scopes to remove from this authorization. |
| scopes | A list of scopes that this authorization is in. |

# orgs


## orgs block-user

https://developer.github.com/v3/orgs/blocking/#block-a-user-from-an-organization



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs check-blocked-user

https://developer.github.com/v3/orgs/blocking/#check-if-a-user-is-blocked-by-an-organization



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs check-membership-for-user

https://developer.github.com/v3/orgs/members/#check-organization-membership-for-a-user

Check if a user is, publicly or privately, a member of the organization.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs check-public-membership-for-user

https://developer.github.com/v3/orgs/members/#check-public-organization-membership-for-a-user



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs convert-member-to-outside-collaborator

https://developer.github.com/v3/orgs/outside_collaborators/#convert-an-organization-member-to-outside-collaborator

When an organization member is converted to an outside collaborator, they'll only have access to the repositories that their current team membership allows. The user will no longer be a member of the organization. For more information, see "[Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)".

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs create-invitation

https://developer.github.com/v3/orgs/members/#create-an-organization-invitation

Invite people to an organization by using their GitHub user ID or their email address. In order to create invitations in an organization, the authenticated user must be an organization owner.

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| email | **Required unless you provide `invitee_id`**. Email address of the person you are inviting, which can be an existing GitHub user. |
| invitee_id | **Required unless you provide `email`**. GitHub user ID for the person you are inviting. |
| role | Specify role for new member. Can be one of:  <br>\* `admin` - Organization owners with full administrative rights to the organization and complete access to all repositories and teams.  <br>\* `direct_member` - Non-owner organization members with ability to see other members and join teams by invitation.  <br>\* `billing_manager` - Non-owner organization members with ability to manage the billing settings of your organization. |
| team_ids | Specify IDs for the teams you want to invite new members to. |

## orgs create-webhook

https://developer.github.com/v3/orgs/hooks/#create-an-organization-webhook

Here's how you can create a hook that posts payloads in JSON format:

### parameters


| name | description |
|------|-------------|
| config.url | __Required__ The URL to which the payloads will be delivered. |
| name | __Required__ Must be passed as "web". |
| org | __Required__  |
| active | Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications. |
| config.content_type | The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`. |
| config.insecure_ssl | Determines whether the SSL certificate of the host for `url` will be verified when delivering payloads. Supported values include `0` (verification is performed) and `1` (verification is not performed). The default is `0`. **We strongly recommend not setting this to `1` as you are subject to man-in-the-middle and other attacks.** |
| config.password |  |
| config.secret | If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value in the [`X-Hub-Signature`](https://developer.github.com/webhooks/event-payloads/#delivery-headers) header. |
| config.username |  |
| events | Determines what [events](https://developer.github.com/webhooks/event-payloads) the hook is triggered for. |

## orgs delete-webhook

https://developer.github.com/v3/orgs/hooks/#delete-an-organization-webhook



### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| org | __Required__  |

## orgs get

https://developer.github.com/v3/orgs/#get-an-organization

To see many of the organization response values, you need to be an authenticated organization owner with the `admin:org` scope. When the value of `two_factor_requirement_enabled` is `true`, the organization requires all members, billing managers, and outside collaborators to enable [two-factor authentication](https://help.github.com/articles/securing-your-account-with-two-factor-authentication-2fa/).

GitHub Apps with the `Organization plan` permission can use this endpoint to retrieve information about an organization's GitHub plan. See "[Authenticating with GitHub Apps](https://developer.github.com/apps/building-github-apps/authenticating-with-github-apps/)" for details. For an example response, see "[Response with GitHub plan information](https://developer.github.com/v3/orgs/#response-with-github-plan-information)."

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| surtur-preview | New repository creation permissions are available to preview. You can now use `members_can_create_public_repositories`, `members_can_create_private_repositories`, and `members_can_create_internal_repositories`. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. These parameters provide more granular permissions to configure the type of repositories organization members can create.<br><br>To access these new parameters during the preview period, you must set this flag. |

## orgs get-membership-for-authenticated-user

https://developer.github.com/v3/orgs/members/#get-an-organization-membership-for-the-authenticated-user



### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## orgs get-membership-for-user

https://developer.github.com/v3/orgs/members/#get-organization-membership-for-a-user

In order to get a user's membership with an organization, the authenticated user must be an organization member.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs get-webhook

https://developer.github.com/v3/orgs/hooks/#get-an-organization-webhook



### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| org | __Required__  |

## orgs list

https://developer.github.com/v3/orgs/#list-organizations

Lists all organizations, in the order that they were created on GitHub.

**Note:** Pagination is powered exclusively by the `since` parameter. Use the [Link header](https://developer.github.com/v3/#link-header) to get the URL for the next page of organizations.

### parameters


| name | description |
|------|-------------|
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## orgs list-app-installations

https://developer.github.com/v3/orgs/#list-app-installations-for-an-organization

Lists all GitHub Apps in an organization. The installation count includes all GitHub Apps installed on repositories in the organization. You must be an organization owner with `admin:read` scope to use this endpoint.

### parameters


| name | description |
|------|-------------|
| machine-man-preview | __Required__ To access the API with your GitHub App, you must set this flag for your requests. |
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## orgs list-blocked-users

https://developer.github.com/v3/orgs/blocking/#list-users-blocked-by-an-organization

List the users blocked by an organization.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## orgs list-for-authenticated-user

https://developer.github.com/v3/orgs/#list-organizations-for-the-authenticated-user

List organizations for the authenticated user.

**OAuth scope requirements**

This only lists organizations that your authorization allows you to operate on in some way (e.g., you can list teams with `read:org` scope, you can publicize your organization membership with `user` scope, etc.). Therefore, this API requires at least `user` or `read:org` scope. OAuth requests with insufficient scope receive a `403 Forbidden` response.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## orgs list-for-user

https://developer.github.com/v3/orgs/#list-organizations-for-a-user

List [public organization memberships](https://help.github.com/articles/publicizing-or-concealing-organization-membership) for the specified user.

This method only lists _public_ memberships, regardless of authentication. If you need to fetch all of the organization memberships (public and private) for the authenticated user, use the [List organizations for the authenticated user](https://developer.github.com/v3/orgs/#list-organizations-for-the-authenticated-user) API instead.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## orgs list-invitation-teams

https://developer.github.com/v3/orgs/members/#list-organization-invitation-teams

List all teams associated with an invitation. In order to see invitations in an organization, the authenticated user must be an organization owner.

### parameters


| name | description |
|------|-------------|
| invitation_id | __Required__ invitation_id parameter |
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## orgs list-members

https://developer.github.com/v3/orgs/members/#list-organization-members

List all users who are members of an organization. If the authenticated user is also a member of this organization then both concealed and public members will be returned.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| filter | Filter members returned in the list. Can be one of:  <br>\* `2fa_disabled` - Members without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled. Available for organization owners.  <br>\* `all` - All members the authenticated user can see. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| role | Filter members returned by their role. Can be one of:  <br>\* `all` - All members of the organization, regardless of role.  <br>\* `admin` - Organization owners.  <br>\* `member` - Non-owner organization members. |

## orgs list-memberships-for-authenticated-user

https://developer.github.com/v3/orgs/members/#list-organization-memberships-for-the-authenticated-user



### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| state | Indicates the state of the memberships to return. Can be either `active` or `pending`. If not specified, the API returns both active and pending memberships. |

## orgs list-outside-collaborators

https://developer.github.com/v3/orgs/outside_collaborators/#list-outside-collaborators-for-an-organization

List all users who are outside collaborators of an organization.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| filter | Filter the list of outside collaborators. Can be one of:  <br>\* `2fa_disabled`: Outside collaborators without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled.  <br>\* `all`: All outside collaborators. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## orgs list-pending-invitations

https://developer.github.com/v3/orgs/members/#list-pending-organization-invitations

The return hash contains a `role` field which refers to the Organization Invitation role and will be one of the following values: `direct_member`, `admin`, `billing_manager`, `hiring_manager`, or `reinstate`. If the invitee is not a GitHub member, the `login` field in the return hash will be `null`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## orgs list-public-members

https://developer.github.com/v3/orgs/members/#list-public-organization-members

Members of an organization can choose to have their membership publicized or not.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## orgs list-saml-sso-authorizations

https://developer.github.com/v3/orgs/#list-saml-sso-authorizations-for-an-organization

Listing and deleting credential authorizations is available to organizations with GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products).

An authenticated organization owner with the `read:org` scope can list all credential authorizations for an organization that uses SAML single sign-on (SSO). The credentials are either personal access tokens or SSH keys that organization members have authorized for the organization. For more information, see [About authentication with SAML single sign-on](https://help.github.com/en/articles/about-authentication-with-saml-single-sign-on).

### parameters


| name | description |
|------|-------------|
| org | __Required__  |

## orgs list-webhooks

https://developer.github.com/v3/orgs/hooks/#list-organization-webhooks



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## orgs ping-webhook

https://developer.github.com/v3/orgs/hooks/#ping-an-organization-webhook

This will trigger a [ping event](https://developer.github.com/webhooks/#ping-event) to be sent to the hook.

### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| org | __Required__  |

## orgs remove-member

https://developer.github.com/v3/orgs/members/#remove-an-organization-member

Removing a user from this list will remove them from all teams and they will no longer have any access to the organization's repositories.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs remove-membership-for-user

https://developer.github.com/v3/orgs/members/#remove-organization-membership-for-a-user

In order to remove a user's membership with an organization, the authenticated user must be an organization owner.

If the specified user is an active member of the organization, this will remove them from the organization. If the specified user has been invited to the organization, this will cancel their invitation. The specified user will receive an email notification in both cases.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs remove-outside-collaborator

https://developer.github.com/v3/orgs/outside_collaborators/#remove-outside-collaborator-from-an-organization

Removing a user from this list will remove them from all the organization's repositories.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs remove-public-membership-for-authenticated-user

https://developer.github.com/v3/orgs/members/#remove-public-organization-membership-for-the-authenticated-user



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs remove-saml-sso-authorization

https://developer.github.com/v3/orgs/#remove-a-saml-sso-authorization-for-an-organization

Listing and deleting credential authorizations is available to organizations with GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products).

An authenticated organization owner with the `admin:org` scope can remove a credential authorization for an organization that uses SAML SSO. Once you remove someone's credential authorization, they will need to create a new personal access token or SSH key and authorize it for the organization they want to access.

### parameters


| name | description |
|------|-------------|
| credential_id | __Required__ credential_id parameter |
| org | __Required__  |

## orgs set-membership-for-user

https://developer.github.com/v3/orgs/members/#set-organization-membership-for-a-user

Only authenticated organization owners can add a member to the organization or update the member's role.

*   If the authenticated user is _adding_ a member to the organization, the invited user will receive an email inviting them to the organization. The user's [membership status](https://developer.github.com/v3/orgs/members/#get-organization-membership-for-a-user) will be `pending` until they accept the invitation.
    
*   Authenticated users can _update_ a user's membership by passing the `role` parameter. If the authenticated user changes a member's role to `admin`, the affected user will receive an email notifying them that they've been made an organization owner. If the authenticated user changes an owner's role to `member`, no email will be sent.

**Rate limits**

To prevent abuse, the authenticated user is limited to 50 organization invitations per 24 hour period. If the organization is more than one month old or on a paid plan, the limit is 500 invitations per 24 hour period.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |
| role | The role to give the user in the organization. Can be one of:  <br>\* `admin` - The user will become an owner of the organization.  <br>\* `member` - The user will become a non-owner member of the organization. |

## orgs set-public-membership-for-authenticated-user

https://developer.github.com/v3/orgs/members/#set-public-organization-membership-for-the-authenticated-user

The user can publicize their own membership. (A user cannot publicize the membership for another user.)

Note that you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs unblock-user

https://developer.github.com/v3/orgs/blocking/#unblock-a-user-from-an-organization



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| username | __Required__  |

## orgs update

https://developer.github.com/v3/orgs/#update-an-organization

**Parameter Deprecation Notice:** GitHub will replace and discontinue `members_allowed_repository_creation_type` in favor of more granular permissions. The new input parameters are `members_can_create_public_repositories`, `members_can_create_private_repositories` for all organizations and `members_can_create_internal_repositories` for organizations associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes).

Enables an authenticated organization owner with the `admin:org` scope to update the organization's profile and member privileges.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| billing_email | Billing email address. This address is not publicized. |
| blog |  |
| company | The company name. |
| default_repository_permission | Default permission level members have for organization repositories:  <br>\* `read` - can pull, but not push to or administer this repository.  <br>\* `write` - can pull and push, but not administer this repository.  <br>\* `admin` - can pull, push, and administer this repository.  <br>\* `none` - no permissions granted by default. |
| description | The description of the company. |
| email | The publicly visible email address. |
| has_organization_projects | Toggles whether an organization can use organization projects. |
| has_repository_projects | Toggles whether repositories that belong to the organization can use repository projects. |
| location | The location. |
| members_allowed_repository_creation_type | Specifies which types of repositories non-admin organization members can create. Can be one of:  <br>\* `all` - all organization members can create public and private repositories.  <br>\* `private` - members can create private repositories. This option is only available to repositories that are part of an organization on GitHub Enterprise Cloud.  <br>\* `none` - only admin members can create repositories.  <br>**Note:** This parameter is deprecated and will be removed in the future. Its return value ignores internal repositories. Using this parameter overrides values set in `members_can_create_repositories`. See [this note](https://developer.github.com/v3/orgs/#members_can_create_repositories) for details. |
| members_can_create_internal_repositories | Toggles whether organization members can create internal repositories, which are visible to all enterprise members. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. Can be one of:  <br>\* `true` - all organization members can create internal repositories.  <br>\* `false` - only organization owners can create internal repositories.  <br>Default: `true`. For more information, see "[Restricting repository creation in your organization](https://help.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation. |
| members_can_create_private_repositories | Toggles whether organization members can create private repositories, which are visible to organization members with permission. Can be one of:  <br>\* `true` - all organization members can create private repositories.  <br>\* `false` - only organization owners can create private repositories.  <br>Default: `true`. For more information, see "[Restricting repository creation in your organization](https://help.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation. |
| members_can_create_public_repositories | Toggles whether organization members can create public repositories, which are visible to anyone. Can be one of:  <br>\* `true` - all organization members can create public repositories.  <br>\* `false` - only organization owners can create public repositories.  <br>Default: `true`. For more information, see "[Restricting repository creation in your organization](https://help.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation. |
| members_can_create_repositories | Toggles the ability of non-admin organization members to create repositories. Can be one of:  <br>\* `true` - all organization members can create repositories.  <br>\* `false` - only organization owners can create repositories.  <br>Default: `true`  <br>**Note:** A parameter can override this parameter. See `members_allowed_repository_creation_type` in this table for details. **Note:** A parameter can override this parameter. See `members_allowed_repository_creation_type` in this table for details. |
| name | The shorthand name of the company. |
| surtur-preview | New repository creation permissions are available to preview. You can now use `members_can_create_public_repositories`, `members_can_create_private_repositories`, and `members_can_create_internal_repositories`. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. These parameters provide more granular permissions to configure the type of repositories organization members can create.<br><br>To access these new parameters during the preview period, you must set this flag. |
| twitter_username | The Twitter username of the company. |

## orgs update-membership-for-authenticated-user

https://developer.github.com/v3/orgs/members/#update-an-organization-membership-for-the-authenticated-user



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| state | __Required__ The state that the membership should be in. Only `"active"` will be accepted. |

## orgs update-webhook

https://developer.github.com/v3/orgs/hooks/#update-an-organization-webhook



### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| org | __Required__  |
| active | Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications. |
| config.content_type | The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`. |
| config.insecure_ssl | Determines whether the SSL certificate of the host for `url` will be verified when delivering payloads. Supported values include `0` (verification is performed) and `1` (verification is not performed). The default is `0`. **We strongly recommend not setting this to `1` as you are subject to man-in-the-middle and other attacks.** |
| config.secret | If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value in the [`X-Hub-Signature`](https://developer.github.com/webhooks/event-payloads/#delivery-headers) header. |
| config.url | The URL to which the payloads will be delivered. |
| events | Determines what [events](https://developer.github.com/webhooks/event-payloads) the hook is triggered for. |
| name |  |

# projects


## projects add-collaborator

https://developer.github.com/v3/projects/collaborators/#add-project-collaborator

Adds a collaborator to an organization project and sets their permission level. You must be an organization owner or a project `admin` to add a collaborator.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |
| username | __Required__  |
| permission | The permission to grant the collaborator. |

## projects create-card

https://developer.github.com/v3/projects/cards/#create-a-project-card

**Note**: GitHub's REST API v3 considers every pull request an issue, but not every issue is a pull request. For this reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by the `pull_request` key.

Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull request id, use the "[List pull requests](https://developer.github.com/v3/pulls/#list-pull-requests)" endpoint.

### parameters


| name | description |
|------|-------------|
| column_id | __Required__ column_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |

## projects create-column

https://developer.github.com/v3/projects/columns/#create-a-project-column



### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| name | __Required__ Name of the project column |
| project_id | __Required__  |

## projects create-for-authenticated-user

https://developer.github.com/v3/projects/#create-a-user-project



### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| name | __Required__ Name of the project |
| body | Body of the project |

## projects create-for-org

https://developer.github.com/v3/projects/#create-an-organization-project

Creates an organization project board. Returns a `404 Not Found` status if projects are disabled in the organization. If you do not have sufficient privileges to perform this action, a `401 Unauthorized` or `410 Gone` status is returned.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| name | __Required__ The name of the project. |
| org | __Required__  |
| body | The description of the project. |

## projects create-for-repo

https://developer.github.com/v3/projects/#create-a-repository-project

Creates a repository project board. Returns a `404 Not Found` status if projects are disabled in the repository. If you do not have sufficient privileges to perform this action, a `401 Unauthorized` or `410 Gone` status is returned.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| name | __Required__ The name of the project. |
| repo | __Required__ repository in OWNER/REPO form |
| body | The description of the project. |

## projects delete

https://developer.github.com/v3/projects/#delete-a-project

Deletes a project board. Returns a `404 Not Found` status if projects are disabled.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |

## projects delete-card

https://developer.github.com/v3/projects/cards/#delete-a-project-card



### parameters


| name | description |
|------|-------------|
| card_id | __Required__ card_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |

## projects delete-column

https://developer.github.com/v3/projects/columns/#delete-a-project-column



### parameters


| name | description |
|------|-------------|
| column_id | __Required__ column_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |

## projects get

https://developer.github.com/v3/projects/#get-a-project

Gets a project by its `id`. Returns a `404 Not Found` status if projects are disabled. If you do not have sufficient privileges to perform this action, a `401 Unauthorized` or `410 Gone` status is returned.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |

## projects get-card

https://developer.github.com/v3/projects/cards/#get-a-project-card



### parameters


| name | description |
|------|-------------|
| card_id | __Required__ card_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |

## projects get-column

https://developer.github.com/v3/projects/columns/#get-a-project-column



### parameters


| name | description |
|------|-------------|
| column_id | __Required__ column_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |

## projects get-permission-for-user

https://developer.github.com/v3/projects/collaborators/#get-project-permission-for-a-user

Returns the collaborator's permission level for an organization project. Possible values for the `permission` key: `admin`, `write`, `read`, `none`. You must be an organization owner or a project `admin` to review a user's permission level.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |
| username | __Required__  |

## projects list-cards

https://developer.github.com/v3/projects/cards/#list-project-cards



### parameters


| name | description |
|------|-------------|
| column_id | __Required__ column_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| archived_state | Filters the project cards that are returned by the card's state. Can be one of `all`,`archived`, or `not_archived`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## projects list-collaborators

https://developer.github.com/v3/projects/collaborators/#list-project-collaborators

Lists the collaborators for an organization project. For a project, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners. You must be an organization owner or a project `admin` to list collaborators.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |
| affiliation | Filters the collaborators by their affiliation. Can be one of:  <br>\* `outside`: Outside collaborators of a project that are not a member of the project's organization.  <br>\* `direct`: Collaborators with permissions to a project, regardless of organization membership status.  <br>\* `all`: All collaborators the authenticated user can see. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## projects list-columns

https://developer.github.com/v3/projects/columns/#list-project-columns



### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## projects list-for-org

https://developer.github.com/v3/projects/#list-organization-projects

Lists the projects in an organization. Returns a `404 Not Found` status if projects are disabled in the organization. If you do not have sufficient privileges to perform this action, a `401 Unauthorized` or `410 Gone` status is returned.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| state | Indicates the state of the projects to return. Can be either `open`, `closed`, or `all`. |

## projects list-for-repo

https://developer.github.com/v3/projects/#list-repository-projects

Lists the projects in a repository. Returns a `404 Not Found` status if projects are disabled in the repository. If you do not have sufficient privileges to perform this action, a `401 Unauthorized` or `410 Gone` status is returned.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| state | Indicates the state of the projects to return. Can be either `open`, `closed`, or `all`. |

## projects list-for-user

https://developer.github.com/v3/projects/#list-user-projects



### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| state | Indicates the state of the projects to return. Can be either `open`, `closed`, or `all`. |

## projects move-card

https://developer.github.com/v3/projects/cards/#move-a-project-card



### parameters


| name | description |
|------|-------------|
| card_id | __Required__ card_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| position | __Required__ The position of the card in a column |
| column_id | The unique identifier of the column the card should be moved to |

## projects move-column

https://developer.github.com/v3/projects/columns/#move-a-project-column



### parameters


| name | description |
|------|-------------|
| column_id | __Required__ column_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| position | __Required__ The position of the column in a project |

## projects remove-collaborator

https://developer.github.com/v3/projects/collaborators/#remove-project-collaborator

Removes a collaborator from an organization project. You must be an organization owner or a project `admin` to remove a collaborator.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |
| username | __Required__  |

## projects update

https://developer.github.com/v3/projects/#update-a-project

Updates a project board's information. Returns a `404 Not Found` status if projects are disabled. If you do not have sufficient privileges to perform this action, a `401 Unauthorized` or `410 Gone` status is returned.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |
| body | Body of the project |
| name | Name of the project |
| organization_permission | The baseline permission that all organization members have on this project |
| private | Whether or not this project can be seen by everyone. |
| state | State of the project; either 'open' or 'closed' |

## projects update-card

https://developer.github.com/v3/projects/cards/#update-a-project-card



### parameters


| name | description |
|------|-------------|
| card_id | __Required__ card_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| archived | Whether or not the card is archived |
| note | The project card's note |

## projects update-column

https://developer.github.com/v3/projects/columns/#update-a-project-column



### parameters


| name | description |
|------|-------------|
| column_id | __Required__ column_id parameter |
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| name | __Required__ Name of the project column |

# pulls


## pulls check-if-merged

https://developer.github.com/v3/pulls/#check-if-a-pull-request-has-been-merged



### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |

## pulls create

https://developer.github.com/v3/pulls/#create-a-pull-request

Draft pull requests are available in public repositories with GitHub Free and GitHub Free for organizations, GitHub Pro, and legacy per-repository billing plans, and in public and private repositories with GitHub Team and GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

To open or update a pull request in a public repository, you must have write access to the head or the source branch. For organization-owned repositories, you must be a member of the organization that owns the repository to open or update a pull request.

You can create a new pull request.

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| base | __Required__ The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository. |
| head | __Required__ The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace `head` with a user like this: `username:branch`. |
| repo | __Required__ repository in OWNER/REPO form |
| body | The contents of the pull request. |
| draft | Indicates whether the pull request is a draft. See "[Draft Pull Requests](https://help.github.com/en/articles/about-pull-requests#draft-pull-requests)" in the GitHub Help documentation to learn more. |
| issue |  |
| maintainer_can_modify | Indicates whether [maintainers can modify](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request. |
| sailor-v-preview | You can now use the REST API to add a reason when you lock an issue, and you will see lock reasons in responses that include issues or pull requests. You will also see lock reasons in `locked` events. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-01-10-lock-reason-api-preview) for full details. To access this feature, you must set this flag. |
| title | The title of the new pull request. |

## pulls create-reply-for-review-comment

https://developer.github.com/v3/pulls/comments/#create-a-reply-for-a-review-comment

Creates a reply to a review comment for a pull request. For the `comment_id`, provide the ID of the review comment you are replying to. This must be the ID of a _top-level review comment_, not a reply to that comment. Replies to replies are not supported.

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The text of the review comment. |
| comment_id | __Required__ comment_id parameter |
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |

## pulls create-review

https://developer.github.com/v3/pulls/reviews/#create-a-review-for-a-pull-request

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

Pull request reviews created in the `PENDING` state do not include the `submitted_at` property in the response.

**Note:** To comment on a specific line in a file, you need to first determine the _position_ of that line in the diff. The GitHub REST API v3 offers the `application/vnd.github.v3.diff` [media type](https://developer.github.com/v3/media/#commits-commit-comparison-and-pull-requests). To see a pull request diff, add this media type to the `Accept` header of a call to the [single pull request](https://developer.github.com/v3/pulls/#get-a-pull-request) endpoint.

The `position` value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| body | **Required** when using `REQUEST_CHANGES` or `COMMENT` for the `event` parameter. The body text of the pull request review. |
| comments | Use the following table to specify the location, destination, and contents of the draft review comment. |
| commit_id | The SHA of the commit that needs a review. Not using the latest commit SHA may render your review comment outdated if a subsequent commit modifies the line you specify as the `position`. Defaults to the most recent commit in the pull request when you do not specify a value. |
| event | The review action you want to perform. The review actions include: `APPROVE`, `REQUEST_CHANGES`, or `COMMENT`. By leaving this blank, you set the review action state to `PENDING`, which means you will need to [submit the pull request review](https://developer.github.com/v3/pulls/reviews/#submit-a-review-for-a-pull-request) when you are ready. |

## pulls create-review-comment

https://developer.github.com/v3/pulls/comments/#create-a-review-comment-for-a-pull-request

**Note:** Multi-line comments on pull requests are currently in public beta and subject to change.

Creates a review comment in the pull request diff. To add a regular comment to a pull request timeline, see "[Create an issue comment](https://developer.github.com/v3/issues/comments/#create-an-issue-comment)." We recommend creating a review comment using `line`, `side`, and optionally `start_line` and `start_side` if your comment applies to more than one line in the pull request diff.

You can still create a review comment using the `position` parameter. When you use `position`, the `line`, `side`, `start_line`, and `start_side` parameters are not required. For more information, see [Multi-line comment summary](https://developer.github.com/v3/pulls/comments/#multi-line-comment-summary-3).

**Note:** The position value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

**Multi-line comment summary**

**Note:** New parameters and response fields are available for developers to preview. During the preview period, these response fields may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2019-10-03-multi-line-comments) for full details.

Use the `comfort-fade` preview header and the `line` parameter to show multi-line comment-supported fields in the response.

If you use the `comfort-fade` preview header, your response will show:

*   For multi-line comments, values for `start_line`, `original_start_line`, `start_side`, `line`, `original_line`, and `side`.
*   For single-line comments, values for `line`, `original_line`, and `side` and a `null` value for `start_line`, `original_start_line`, and `start_side`.

If you don't use the `comfort-fade` preview header, multi-line and single-line comments will appear the same way in the response with a single `position` attribute. Your response will show:

*   For multi-line comments, the last line of the comment range for the `position` attribute.
*   For single-line comments, the diff-positioned way of referencing comments for the `position` attribute. For more information, see `position` in the [input parameters](https://developer.github.com/v3/pulls/comments/#parameters-2) table.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The text of the review comment. |
| path | __Required__ The relative path to the file that necessitates a comment. |
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| comfort-fade-preview | Multi-line comments in a pull request diff is currently available for developers to preview. To access the new response fields during the preview period, you must set this flag. |
| commit_id | The SHA of the commit needing a comment. Not using the latest commit SHA may render your comment outdated if a subsequent commit modifies the line you specify as the `position`. |
| in_reply_to |  |
| line | **Required with `comfort-fade` preview**. The line of the blob in the pull request diff that the comment applies to. For a multi-line comment, the last line of the range that your comment applies to. |
| position | **Required without `comfort-fade` preview**. The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. For help finding the position value, read the note above. |
| side | **Required with `comfort-fade` preview**. In a split diff view, the side of the diff that the pull request's changes appear on. Can be `LEFT` or `RIGHT`. Use `LEFT` for deletions that appear in red. Use `RIGHT` for additions that appear in green or unchanged lines that appear in white and are shown for context. For a multi-line comment, side represents whether the last line of the comment range is a deletion or addition. For more information, see "[Diff view options](https://help.github.com/en/articles/about-comparing-branches-in-pull-requests#diff-view-options)" in the GitHub Help documentation. |
| start_line | **Required when using multi-line comments**. To create multi-line comments, you must use the `comfort-fade` preview header. The `start_line` is the first line in the pull request diff that your multi-line comment applies to. To learn more about multi-line comments, see "[Commenting on a pull request](https://help.github.com/en/articles/commenting-on-a-pull-request#adding-line-comments-to-a-pull-request)" in the GitHub Help documentation. |
| start_side | **Required when using multi-line comments**. To create multi-line comments, you must use the `comfort-fade` preview header. The `start_side` is the starting side of the diff that the comment applies to. Can be `LEFT` or `RIGHT`. To learn more about multi-line comments, see "[Commenting on a pull request](https://help.github.com/en/articles/commenting-on-a-pull-request#adding-line-comments-to-a-pull-request)" in the GitHub Help documentation. See `side` in this table for additional context. |

## pulls delete-pending-review

https://developer.github.com/v3/pulls/reviews/#delete-a-pending-review-for-a-pull-request



### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| review_id | __Required__ review_id parameter |

## pulls delete-review-comment

https://developer.github.com/v3/pulls/comments/#delete-a-review-comment-for-a-pull-request

Deletes a review comment.

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## pulls dismiss-review

https://developer.github.com/v3/pulls/reviews/#dismiss-a-review-for-a-pull-request

**Note:** To dismiss a pull request review on a [protected branch](https://developer.github.com/v3/repos/branches/), you must be a repository administrator or be included in the list of people or teams who can dismiss pull request reviews.

### parameters


| name | description |
|------|-------------|
| message | __Required__ The message for the pull request review dismissal |
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| review_id | __Required__ review_id parameter |
| event |  |

## pulls get

https://developer.github.com/v3/pulls/#get-a-pull-request

Draft pull requests are available in public repositories with GitHub Free and GitHub Free for organizations, GitHub Pro, and legacy per-repository billing plans, and in public and private repositories with GitHub Team and GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Lists details of a pull request by providing its number.

When you get, [create](https://developer.github.com/v3/pulls/#create-a-pull-request), or [edit](https://developer.github.com/v3/pulls/#update-a-pull-request) a pull request, GitHub creates a merge commit to test whether the pull request can be automatically merged into the base branch. This test commit is not added to the base branch or the head branch. You can review the status of the test commit using the `mergeable` key. For more information, see "[Checking mergeability of pull requests](https://developer.github.com/v3/git/#checking-mergeability-of-pull-requests)".

The value of the `mergeable` attribute can be `true`, `false`, or `null`. If the value is `null`, then GitHub has started a background job to compute the mergeability. After giving the job time to complete, resubmit the request. When the job finishes, you will see a non-`null` value for the `mergeable` attribute in the response. If `mergeable` is `true`, then `merge_commit_sha` will be the SHA of the _test_ merge commit.

The value of the `merge_commit_sha` attribute changes depending on the state of the pull request. Before merging a pull request, the `merge_commit_sha` attribute holds the SHA of the _test_ merge commit. After merging a pull request, the `merge_commit_sha` attribute changes depending on how you merged the pull request:

*   If merged as a [merge commit](https://help.github.com/articles/about-merge-methods-on-github/), `merge_commit_sha` represents the SHA of the merge commit.
*   If merged via a [squash](https://help.github.com/articles/about-merge-methods-on-github/#squashing-your-merge-commits), `merge_commit_sha` represents the SHA of the squashed commit on the base branch.
*   If [rebased](https://help.github.com/articles/about-merge-methods-on-github/#rebasing-and-merging-your-commits), `merge_commit_sha` represents the commit that the base branch was updated to.

Pass the appropriate [media type](https://developer.github.com/v3/media/#commits-commit-comparison-and-pull-requests) to fetch diff and patch formats.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| sailor-v-preview | You can now use the REST API to add a reason when you lock an issue, and you will see lock reasons in responses that include issues or pull requests. You will also see lock reasons in `locked` events. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-01-10-lock-reason-api-preview) for full details. To access this feature, you must set this flag. |

## pulls get-review

https://developer.github.com/v3/pulls/reviews/#get-a-review-for-a-pull-request



### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| review_id | __Required__ review_id parameter |

## pulls get-review-comment

https://developer.github.com/v3/pulls/comments/#get-a-review-comment-for-a-pull-request

**Note:** Multi-line comments on pull requests are currently in public beta and subject to change.

Provides details for a review comment.

**Multi-line comment summary**

**Note:** New parameters and response fields are available for developers to preview. During the preview period, these response fields may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2019-10-03-multi-line-comments) for full details.

Use the `comfort-fade` preview header and the `line` parameter to show multi-line comment-supported fields in the response.

If you use the `comfort-fade` preview header, your response will show:

*   For multi-line comments, values for `start_line`, `original_start_line`, `start_side`, `line`, `original_line`, and `side`.
*   For single-line comments, values for `line`, `original_line`, and `side` and a `null` value for `start_line`, `original_start_line`, and `start_side`.

If you don't use the `comfort-fade` preview header, multi-line and single-line comments will appear the same way in the response with a single `position` attribute. Your response will show:

*   For multi-line comments, the last line of the comment range for the `position` attribute.
*   For single-line comments, the diff-positioned way of referencing comments for the `position` attribute. For more information, see `position` in the [input parameters](https://developer.github.com/v3/pulls/comments/#parameters-2) table.

The `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](https://developer.github.com/v3/reactions) reactions.

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| comfort-fade-preview | Multi-line comments in a pull request diff is currently available for developers to preview. To access the new response fields during the preview period, you must set this flag. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## pulls list

https://developer.github.com/v3/pulls/#list-pull-requests

Draft pull requests are available in public repositories with GitHub Free and GitHub Free for organizations, GitHub Pro, and legacy per-repository billing plans, and in public and private repositories with GitHub Team and GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| base | Filter pulls by base branch name. Example: `gh-pages`. |
| direction | The direction of the sort. Can be either `asc` or `desc`. Default: `desc` when sort is `created` or sort is not specified, otherwise `asc`. |
| head | Filter pulls by head user or head organization and branch name in the format of `user:ref-name` or `organization:ref-name`. For example: `github:new-script-format` or `octocat:test-branch`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sailor-v-preview | You can now use the REST API to add a reason when you lock an issue, and you will see lock reasons in responses that include issues or pull requests. You will also see lock reasons in `locked` events. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-01-10-lock-reason-api-preview) for full details. To access this feature, you must set this flag. |
| sort | What to sort results by. Can be either `created`, `updated`, `popularity` (comment count) or `long-running` (age, filtering by pulls updated in the last month). |
| state | Either `open`, `closed`, or `all` to filter by state. |

## pulls list-comments-for-review

https://developer.github.com/v3/pulls/reviews/#list-comments-for-a-pull-request-review

List comments for a specific pull request review.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| review_id | __Required__ review_id parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## pulls list-commits

https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request

Lists a maximum of 250 commits for a pull request. To receive a complete commit list for pull requests with more than 250 commits, use the [List commits](https://developer.github.com/v3/repos/commits/#list-commits) endpoint.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## pulls list-files

https://developer.github.com/v3/pulls/#list-pull-requests-files

**Note:** Responses include a maximum of 3000 files. The paginated response returns 30 files per page by default.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## pulls list-requested-reviewers

https://developer.github.com/v3/pulls/review_requests/#list-requested-reviewers-for-a-pull-request



### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## pulls list-review-comments

https://developer.github.com/v3/pulls/comments/#list-review-comments-on-a-pull-request

**Note:** Multi-line comments on pull requests are currently in public beta and subject to change.

Lists all review comments for a pull request. By default, review comments are in ascending order by ID.

**Multi-line comment summary**

**Note:** New parameters and response fields are available for developers to preview. During the preview period, these response fields may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2019-10-03-multi-line-comments) for full details.

Use the `comfort-fade` preview header and the `line` parameter to show multi-line comment-supported fields in the response.

If you use the `comfort-fade` preview header, your response will show:

*   For multi-line comments, values for `start_line`, `original_start_line`, `start_side`, `line`, `original_line`, and `side`.
*   For single-line comments, values for `line`, `original_line`, and `side` and a `null` value for `start_line`, `original_start_line`, and `start_side`.

If you don't use the `comfort-fade` preview header, multi-line and single-line comments will appear the same way in the response with a single `position` attribute. Your response will show:

*   For multi-line comments, the last line of the comment range for the `position` attribute.
*   For single-line comments, the diff-positioned way of referencing comments for the `position` attribute. For more information, see `position` in the [input parameters](https://developer.github.com/v3/pulls/comments/#parameters-2) table.

The `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](https://developer.github.com/v3/reactions) reactions.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| comfort-fade-preview | Multi-line comments in a pull request diff is currently available for developers to preview. To access the new response fields during the preview period, you must set this flag. |
| direction | Can be either `asc` or `desc`. Ignored without `sort` parameter. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| sort | One of `created` (when the repository was starred) or `updated` (when it was last pushed to). |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## pulls list-review-comments-for-repo

https://developer.github.com/v3/pulls/comments/#list-review-comments-in-a-repository

**Note:** Multi-line comments on pull requests are currently in public beta and subject to change.

Lists review comments for all pull requests in a repository. By default, review comments are in ascending order by ID.

**Multi-line comment summary**

**Note:** New parameters and response fields are available for developers to preview. During the preview period, these response fields may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2019-10-03-multi-line-comments) for full details.

Use the `comfort-fade` preview header and the `line` parameter to show multi-line comment-supported fields in the response.

If you use the `comfort-fade` preview header, your response will show:

*   For multi-line comments, values for `start_line`, `original_start_line`, `start_side`, `line`, `original_line`, and `side`.
*   For single-line comments, values for `line`, `original_line`, and `side` and a `null` value for `start_line`, `original_start_line`, and `start_side`.

If you don't use the `comfort-fade` preview header, multi-line and single-line comments will appear the same way in the response with a single `position` attribute. Your response will show:

*   For multi-line comments, the last line of the comment range for the `position` attribute.
*   For single-line comments, the diff-positioned way of referencing comments for the `position` attribute. For more information, see `position` in the [input parameters](https://developer.github.com/v3/pulls/comments/#parameters-2) table.

The `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](https://developer.github.com/v3/reactions) reactions.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| comfort-fade-preview | Multi-line comments in a pull request diff is currently available for developers to preview. To access the new response fields during the preview period, you must set this flag. |
| direction | Can be either `asc` or `desc`. Ignored without `sort` parameter. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| sort | One of `created` (when the repository was starred) or `updated` (when it was last pushed to). |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## pulls list-reviews

https://developer.github.com/v3/pulls/reviews/#list-reviews-for-a-pull-request

The list of reviews returns in chronological order.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## pulls merge

https://developer.github.com/v3/pulls/#merge-a-pull-request

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| commit_message | Extra detail to append to automatic commit message. |
| commit_title | Title for the automatic commit message. |
| merge_method | Merge method to use. Possible values are `merge`, `squash` or `rebase`. Default is `merge`. |
| sha | SHA that pull request head must match to allow merge. |

## pulls remove-requested-reviewers

https://developer.github.com/v3/pulls/review_requests/#remove-requested-reviewers-from-a-pull-request



### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| reviewers | An array of user `login`s that will be removed. |
| team_reviewers | An array of team `slug`s that will be removed. |

## pulls request-reviewers

https://developer.github.com/v3/pulls/review_requests/#request-reviewers-for-a-pull-request

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| reviewers | An array of user `login`s that will be requested. |
| team_reviewers | An array of team `slug`s that will be requested. |

## pulls submit-review

https://developer.github.com/v3/pulls/reviews/#submit-a-review-for-a-pull-request



### parameters


| name | description |
|------|-------------|
| event | __Required__ The review action you want to perform. The review actions include: `APPROVE`, `REQUEST_CHANGES`, or `COMMENT`. When you leave this blank, the API returns _HTTP 422 (Unrecognizable entity)_ and sets the review action state to `PENDING`, which means you will need to re-submit the pull request review using a review action. |
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| review_id | __Required__ review_id parameter |
| body | The body text of the pull request review |

## pulls update

https://developer.github.com/v3/pulls/#update-a-pull-request

Draft pull requests are available in public repositories with GitHub Free and GitHub Free for organizations, GitHub Pro, and legacy per-repository billing plans, and in public and private repositories with GitHub Team and GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

To open or update a pull request in a public repository, you must have write access to the head or the source branch. For organization-owned repositories, you must be a member of the organization that owns the repository to open or update a pull request.

### parameters


| name | description |
|------|-------------|
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| base | The name of the branch you want your changes pulled into. This should be an existing branch on the current repository. You cannot update the base branch on a pull request to point to another repository. |
| body | The contents of the pull request. |
| maintainer_can_modify | Indicates whether [maintainers can modify](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request. |
| sailor-v-preview | You can now use the REST API to add a reason when you lock an issue, and you will see lock reasons in responses that include issues or pull requests. You will also see lock reasons in `locked` events. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-01-10-lock-reason-api-preview) for full details. To access this feature, you must set this flag. |
| state | State of this Pull Request. Either `open` or `closed`. |
| title | The title of the pull request. |

## pulls update-branch

https://developer.github.com/v3/pulls/#update-a-pull-request-branch

Updates the pull request branch with the latest upstream changes by merging HEAD from the base branch into the pull request branch.

### parameters


| name | description |
|------|-------------|
| lydian-preview | __Required__ Updating the pull request branch with latest upstream changes is currently available for developers to preview. To access this new endpoint during the preview period, you must set this flag. |
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| expected_head_sha | The expected SHA of the pull request's HEAD ref. This is the most recent commit on the pull request's branch. If the expected SHA does not match the pull request's HEAD, you will receive a `422 Unprocessable Entity` status. You can use the "[List commits](https://developer.github.com/v3/repos/commits/#list-commits)" endpoint to find the most recent commit SHA. Default: SHA of the pull request's current HEAD ref. |

## pulls update-review

https://developer.github.com/v3/pulls/reviews/#update-a-review-for-a-pull-request

Update the review summary comment with new text.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The body text of the pull request review. |
| pull_number | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| review_id | __Required__ review_id parameter |

## pulls update-review-comment

https://developer.github.com/v3/pulls/comments/#update-a-review-comment-for-a-pull-request

**Note:** Multi-line comments on pull requests are currently in public beta and subject to change.

Enables you to edit a review comment.

**Multi-line comment summary**

**Note:** New parameters and response fields are available for developers to preview. During the preview period, these response fields may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2019-10-03-multi-line-comments) for full details.

Use the `comfort-fade` preview header and the `line` parameter to show multi-line comment-supported fields in the response.

If you use the `comfort-fade` preview header, your response will show:

*   For multi-line comments, values for `start_line`, `original_start_line`, `start_side`, `line`, `original_line`, and `side`.
*   For single-line comments, values for `line`, `original_line`, and `side` and a `null` value for `start_line`, `original_start_line`, and `start_side`.

If you don't use the `comfort-fade` preview header, multi-line and single-line comments will appear the same way in the response with a single `position` attribute. Your response will show:

*   For multi-line comments, the last line of the comment range for the `position` attribute.
*   For single-line comments, the diff-positioned way of referencing comments for the `position` attribute. For more information, see `position` in the [input parameters](https://developer.github.com/v3/pulls/comments/#parameters-2) table.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The text of the reply to the review comment. |
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| comfort-fade-preview | Multi-line comments in a pull request diff is currently available for developers to preview. To access the new response fields during the preview period, you must set this flag. |

# rate-limit


## rate-limit get

https://developer.github.com/v3/rate_limit/#get-rate-limit-status-for-the-authenticated-user

**Note:** Accessing this endpoint does not count against your REST API rate limit.

**Note:** The `rate` object is deprecated. If you're writing new API client code or updating existing code, you should use the `core` object instead of the `rate` object. The `core` object contains the same information that is present in the `rate` object.

# reactions


## reactions create-for-commit-comment

https://developer.github.com/v3/reactions/#create-reaction-for-a-commit-comment

Create a reaction to a [commit comment](https://developer.github.com/v3/repos/comments/). A response with a `Status: 200 OK` means that you already added the reaction type to this commit comment.

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| content | __Required__ The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the commit comment. |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions create-for-issue

https://developer.github.com/v3/reactions/#create-reaction-for-an-issue

Create a reaction to an [issue](https://developer.github.com/v3/issues/). A response with a `Status: 200 OK` means that you already added the reaction type to this issue.

### parameters


| name | description |
|------|-------------|
| content | __Required__ The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the issue. |
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions create-for-issue-comment

https://developer.github.com/v3/reactions/#create-reaction-for-an-issue-comment

Create a reaction to an [issue comment](https://developer.github.com/v3/issues/comments/). A response with a `Status: 200 OK` means that you already added the reaction type to this issue comment.

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| content | __Required__ The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the issue comment. |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions create-for-pull-request-review-comment

https://developer.github.com/v3/reactions/#create-reaction-for-a-pull-request-review-comment

Create a reaction to a [pull request review comment](https://developer.github.com/v3/pulls/comments/). A response with a `Status: 200 OK` means that you already added the reaction type to this pull request review comment.

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| content | __Required__ The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the pull request review comment. |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions create-for-team-discussion-comment-in-org

https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-comment

Create a reaction to a [team discussion comment](https://developer.github.com/v3/teams/discussion_comments/). OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/). A response with a `Status: 200 OK` means that you already added the reaction type to this team discussion comment.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `POST /organizations/:org_id/team/:team_id/discussions/:discussion_number/comments/:comment_number/reactions`.

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| content | __Required__ The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the team discussion comment. |
| discussion_number | __Required__  |
| org | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_slug | __Required__ team_slug parameter |

## reactions create-for-team-discussion-comment-legacy

https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-comment-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Create reaction for a team discussion comment`](https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-comment) endpoint.

Create a reaction to a [team discussion comment](https://developer.github.com/v3/teams/discussion_comments/). OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/). A response with a `Status: 200 OK` means that you already added the reaction type to this team discussion comment.

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| content | __Required__ The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the team discussion comment. |
| discussion_number | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_id | __Required__  |

## reactions create-for-team-discussion-in-org

https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion

Create a reaction to a [team discussion](https://developer.github.com/v3/teams/discussions/). OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/). A response with a `Status: 200 OK` means that you already added the reaction type to this team discussion.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `POST /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.

### parameters


| name | description |
|------|-------------|
| content | __Required__ The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the team discussion. |
| discussion_number | __Required__  |
| org | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_slug | __Required__ team_slug parameter |

## reactions create-for-team-discussion-legacy

https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Create reaction for a team discussion`](https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion) endpoint.

Create a reaction to a [team discussion](https://developer.github.com/v3/teams/discussions/). OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/). A response with a `Status: 200 OK` means that you already added the reaction type to this team discussion.

### parameters


| name | description |
|------|-------------|
| content | __Required__ The [reaction type](https://developer.github.com/v3/reactions/#reaction-types) to add to the team discussion. |
| discussion_number | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_id | __Required__  |

## reactions delete-for-commit-comment

https://developer.github.com/v3/reactions/#delete-a-commit-comment-reaction

**Note:** You can also specify a repository by `repository_id` using the route `DELETE /repositories/:repository_id/comments/:comment_id/reactions/:reaction_id`.

Delete a reaction to a [commit comment](https://developer.github.com/v3/repos/comments/).

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| reaction_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions delete-for-issue

https://developer.github.com/v3/reactions/#delete-an-issue-reaction

**Note:** You can also specify a repository by `repository_id` using the route `DELETE /repositories/:repository_id/issues/:issue_number/reactions/:reaction_id`.

Delete a reaction to an [issue](https://developer.github.com/v3/issues/).

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| reaction_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions delete-for-issue-comment

https://developer.github.com/v3/reactions/#delete-an-issue-comment-reaction

**Note:** You can also specify a repository by `repository_id` using the route `DELETE delete /repositories/:repository_id/issues/comments/:comment_id/reactions/:reaction_id`.

Delete a reaction to an [issue comment](https://developer.github.com/v3/issues/comments/).

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| reaction_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions delete-for-pull-request-comment

https://developer.github.com/v3/reactions/#delete-a-pull-request-comment-reaction

**Note:** You can also specify a repository by `repository_id` using the route `DELETE /repositories/:repository_id/pulls/comments/:comment_id/reactions/:reaction_id.`

Delete a reaction to a [pull request review comment](https://developer.github.com/v3/pulls/comments/).

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| reaction_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions delete-for-team-discussion

https://developer.github.com/v3/reactions/#delete-team-discussion-reaction

**Note:** You can also specify a team or organization with `team_id` and `org_id` using the route `DELETE /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions/:reaction_id`.

Delete a reaction to a [team discussion](https://developer.github.com/v3/teams/discussions/). OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| org | __Required__  |
| reaction_id | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_slug | __Required__ team_slug parameter |

## reactions delete-for-team-discussion-comment

https://developer.github.com/v3/reactions/#delete-team-discussion-comment-reaction

**Note:** You can also specify a team or organization with `team_id` and `org_id` using the route `DELETE /organizations/:org_id/team/:team_id/discussions/:discussion_number/comments/:comment_number/reactions/:reaction_id`.

Delete a reaction to a [team discussion comment](https://developer.github.com/v3/teams/discussion_comments/). OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| discussion_number | __Required__  |
| org | __Required__  |
| reaction_id | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_slug | __Required__ team_slug parameter |

## reactions delete-legacy

https://developer.github.com/v3/reactions/#delete-a-reaction-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Reactions API. We recommend migrating your existing code to use the new delete reactions endpoints. For more information, see this [blog post](https://developer.github.com/changes/2020-02-26-new-delete-reactions-endpoints/).

OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/), when deleting a [team discussion](https://developer.github.com/v3/teams/discussions/) or [team discussion comment](https://developer.github.com/v3/teams/discussion_comments/).

### parameters


| name | description |
|------|-------------|
| reaction_id | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## reactions list-for-commit-comment

https://developer.github.com/v3/reactions/#list-reactions-for-a-commit-comment

List the reactions to a [commit comment](https://developer.github.com/v3/repos/comments/).

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| content | Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a commit comment. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## reactions list-for-issue

https://developer.github.com/v3/reactions/#list-reactions-for-an-issue

List the reactions to an [issue](https://developer.github.com/v3/issues/).

### parameters


| name | description |
|------|-------------|
| issue_number | __Required__ issue_number parameter |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| content | Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to an issue. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## reactions list-for-issue-comment

https://developer.github.com/v3/reactions/#list-reactions-for-an-issue-comment

List the reactions to an [issue comment](https://developer.github.com/v3/issues/comments/).

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| content | Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to an issue comment. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## reactions list-for-pull-request-review-comment

https://developer.github.com/v3/reactions/#list-reactions-for-a-pull-request-review-comment

List the reactions to a [pull request review comment](https://developer.github.com/v3/pulls/comments/).

### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| content | Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a pull request review comment. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## reactions list-for-team-discussion-comment-in-org

https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-comment

List the reactions to a [team discussion comment](https://developer.github.com/v3/teams/discussion_comments/). OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/:org_id/team/:team_id/discussions/:discussion_number/comments/:comment_number/reactions`.

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| discussion_number | __Required__  |
| org | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_slug | __Required__ team_slug parameter |
| content | Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a team discussion comment. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## reactions list-for-team-discussion-comment-legacy

https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-comment-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion comment`](https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-comment) endpoint.

List the reactions to a [team discussion comment](https://developer.github.com/v3/teams/discussion_comments/). OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| discussion_number | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_id | __Required__  |
| content | Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a team discussion comment. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## reactions list-for-team-discussion-in-org

https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion

List the reactions to a [team discussion](https://developer.github.com/v3/teams/discussions/). OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/:org_id/team/:team_id/discussions/:discussion_number/reactions`.

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| org | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_slug | __Required__ team_slug parameter |
| content | Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a team discussion. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## reactions list-for-team-discussion-legacy

https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion`](https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion) endpoint.

List the reactions to a [team discussion](https://developer.github.com/v3/teams/discussions/). OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| squirrel-girl-preview | __Required__ An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| team_id | __Required__  |
| content | Returns a single [reaction type](https://developer.github.com/v3/reactions/#reaction-types). Omit this parameter to list all reactions to a team discussion. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

# repos


## repos accept-invitation

https://developer.github.com/v3/repos/invitations/#accept-a-repository-invitation



### parameters


| name | description |
|------|-------------|
| invitation_id | __Required__ invitation_id parameter |

## repos add-app-access-restrictions

https://developer.github.com/v3/repos/branches/#add-app-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Grants the specified apps push access for this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.

| Type    | Description                                                                                                                                                |
| ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `array` | The GitHub Apps that have push access to this branch. Use the app's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos add-collaborator

https://developer.github.com/v3/repos/collaborators/#add-a-repository-collaborator

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

For more information the permission levels, see "[Repository permission levels for an organization](https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/repository-permission-levels-for-an-organization#permission-levels-for-repositories-owned-by-an-organization)".

Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

The invitee will receive a notification that they have been invited to the repository, which they must accept or decline. They may do this via the notifications page, the email they receive, or by using the [repository invitations API endpoints](https://developer.github.com/v3/repos/invitations/).

**Rate limits**

To prevent abuse, you are limited to sending 50 invitations to a repository per 24 hour period. Note there is no limit if you are inviting organization members to an organization repository.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| username | __Required__  |
| permission | The permission to grant the collaborator. **Only valid on organization-owned repositories.** Can be one of:  <br>\* `pull` - can pull, but not push to or administer this repository.  <br>\* `push` - can pull and push, but not administer this repository.  <br>\* `admin` - can pull, push and administer this repository.  <br>\* `maintain` - Recommended for project managers who need to manage the repository without access to sensitive or destructive actions.  <br>\* `triage` - Recommended for contributors who need to proactively manage issues and pull requests without write access. |
| permissions |  |

## repos add-status-check-contexts

https://developer.github.com/v3/repos/branches/#add-status-check-contexts

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos add-team-access-restrictions

https://developer.github.com/v3/repos/branches/#add-team-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Grants the specified teams push access for this branch. You can also give push access to child teams.

| Type    | Description                                                                                                                                |
| ------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `array` | The teams that can have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos add-user-access-restrictions

https://developer.github.com/v3/repos/branches/#add-user-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Grants the specified people push access for this branch.

| Type    | Description                                                                                                                   |
| ------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos check-collaborator

https://developer.github.com/v3/repos/collaborators/#check-if-a-user-is-a-repository-collaborator

For organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.

Team members will include the members of child teams.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| username | __Required__  |

## repos check-vulnerability-alerts

https://developer.github.com/v3/repos/#check-if-vulnerability-alerts-are-enabled-for-a-repository

Shows whether dependency alerts are enabled or disabled for a repository. The authenticated user must have admin access to the repository. For more information, see "[About security alerts for vulnerable dependencies](https://help.github.com/en/articles/about-security-alerts-for-vulnerable-dependencies)".

### parameters


| name | description |
|------|-------------|
| dorian-preview | __Required__ Enabling and disabling dependency alerts for a repository using the REST API is currently available for developers to preview. To access these new endpoints during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |

## repos compare-commits

https://developer.github.com/v3/repos/commits/#compare-two-commits

Both `:base` and `:head` must be branch names in `:repo`. To compare branches across other repositories in the same network as `:repo`, use the format `<USERNAME>:branch`.

The response from the API is equivalent to running the `git log base..head` command; however, commits are returned in chronological order. Pass the appropriate [media type](https://developer.github.com/v3/media/#commits-commit-comparison-and-pull-requests) to fetch diff and patch formats.

The response also includes details on the files that were changed between the two commits. This includes the status of the change (for example, if a file was added, removed, modified, or renamed), and details of the change itself. For example, files with a `renamed` status have a `previous_filename` field showing the previous filename of the file, and files with a `modified` status have a `patch` field showing the changes made to the file.

**Working with large comparisons**

The response will include a comparison of up to 250 commits. If you are working with a larger commit range, you can use the [List commits](https://developer.github.com/v3/repos/commits/#list-commits) to enumerate all commits in the range.

For comparisons with extremely large diffs, you may receive an error response indicating that the diff took too long to generate. You can typically resolve this error by using a smaller commit range.

**Signature verification object**

The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:

These are the possible values for `reason` in the `verification` object:

| Value                    | Description                                                                                                                       |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| `expired_key`            | The key that made the signature is expired.                                                                                       |
| `not_signing_key`        | The "signing" flag is not among the usage flags in the GPG key that made the signature.                                           |
| `gpgverify_error`        | There was an error communicating with the signature verification service.                                                         |
| `gpgverify_unavailable`  | The signature verification service is currently unavailable.                                                                      |
| `unsigned`               | The object does not include a signature.                                                                                          |
| `unknown_signature_type` | A non-PGP signature was found in the commit.                                                                                      |
| `no_user`                | No user was associated with the `committer` email address in the commit.                                                          |
| `unverified_email`       | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. |
| `bad_email`              | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature.             |
| `unknown_key`            | The key that made the signature has not been registered with any user's account.                                                  |
| `malformed_signature`    | There was an error parsing the signature.                                                                                         |
| `invalid`                | The signature could not be cryptographically verified using the key whose key-id was found in the signature.                      |
| `valid`                  | None of the above errors applied, so the signature is considered to be verified.                                                  |

### parameters


| name | description |
|------|-------------|
| base | __Required__ base parameter |
| head | __Required__ head parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos create-commit-comment

https://developer.github.com/v3/repos/comments/#create-a-commit-comment

Create a comment for a commit using its `:commit_sha`.

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The contents of the comment. |
| commit_sha | __Required__ commit_sha+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| line | **Deprecated**. Use **position** parameter instead. Line number in the file to comment on. |
| path | Relative path of the file to comment on. |
| position | Line index in the diff to comment on. |

## repos create-commit-signature-protection

https://developer.github.com/v3/repos/branches/#create-commit-signature-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

When authenticated with admin or owner permissions to the repository, you can use this endpoint to require signed commits on a branch. You must enable branch protection to require signed commits.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| zzzax-preview | __Required__ Protected Branches API can now manage a setting for requiring signed commits. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-02-22-protected-branches-required-signatures) for full details. To access the API during the preview period, you must set this flag. |

## repos create-commit-status

https://developer.github.com/v3/repos/statuses/#create-a-commit-status

Users with push access in a repository can create commit statuses for a given SHA.

Note: there is a limit of 1000 statuses per `sha` and `context` within a repository. Attempts to create more than 1000 statuses will result in a validation error.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| sha | __Required__ sha parameter |
| state | __Required__ The state of the status. Can be one of `error`, `failure`, `pending`, or `success`. |
| context | A string label to differentiate this status from the status of other systems. |
| description | A short description of the status. |
| target_url | The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the source of the status.  <br>For example, if your continuous integration system is posting build status, you would want to provide the deep link for the build output for this specific SHA:  <br>`http://ci.example.com/user/repo/build/sha` |

## repos create-deploy-key

https://developer.github.com/v3/repos/keys/#create-a-deploy-key

You can create a read-only deploy key.

### parameters


| name | description |
|------|-------------|
| key | __Required__ The contents of the key. |
| repo | __Required__ repository in OWNER/REPO form |
| read_only | If `true`, the key will only be able to read repository contents. Otherwise, the key will be able to read and write.  <br>  <br>Deploy keys with write access can perform the same actions as an organization member with admin access, or a collaborator on a personal repository. For more information, see "[Repository permission levels for an organization](https://help.github.com/articles/repository-permission-levels-for-an-organization/)" and "[Permission levels for a user account repository](https://help.github.com/articles/permission-levels-for-a-user-account-repository/)." |
| title | A name for the key. |

## repos create-deployment

https://developer.github.com/v3/repos/deployments/#create-a-deployment

Deployments offer a few configurable parameters with certain defaults.

The `ref` parameter can be any named branch, tag, or SHA. At GitHub we often deploy branches and verify them
before we merge a pull request.

The `environment` parameter allows deployments to be issued to different runtime environments. Teams often have
multiple environments for verifying their applications, such as `production`, `staging`, and `qa`. This parameter
makes it easier to track which environments have requested deployments. The default environment is `production`.

The `auto_merge` parameter is used to ensure that the requested ref is not behind the repository's default branch. If
the ref _is_ behind the default branch for the repository, we will attempt to merge it for you. If the merge succeeds,
the API will return a successful merge commit. If merge conflicts prevent the merge from succeeding, the API will
return a failure response.

By default, [commit statuses](https://developer.github.com/v3/repos/statuses) for every submitted context must be in a `success`
state. The `required_contexts` parameter allows you to specify a subset of contexts that must be `success`, or to
specify contexts that have not yet been submitted. You are not required to use commit statuses to deploy. If you do
not require any contexts or create any commit statuses, the deployment will always succeed.

The `payload` parameter is available for any extra information that a deployment system might need. It is a JSON text
field that will be passed on when a deployment event is dispatched.

The `task` parameter is used by the deployment system to allow different execution paths. In the web world this might
be `deploy:migrations` to run schema changes on the system. In the compiled world this could be a flag to compile an
application with debugging enabled.

Users with `repo` or `repo_deployment` scopes can create a deployment for a given ref.

#### Merged branch response
You will see this response when GitHub automatically merges the base branch into the topic branch instead of creating
a deployment. This auto-merge happens when:
*   Auto-merge option is enabled in the repository
*   Topic branch does not include the latest changes on the base branch, which is `master` in the response example
*   There are no merge conflicts

If there are no new commits in the base branch, a new request to create a deployment should give a successful
response.

#### Merge conflict response
This error happens when the `auto_merge` option is enabled and when the default branch (in this case `master`), can't
be merged into the branch that's being deployed (in this case `topic-branch`), due to merge conflicts.

#### Failed commit status checks
This error happens when the `required_contexts` parameter indicates that one or more contexts need to have a `success`
status for the commit to be deployed, but one or more of the required contexts do not have a state of `success`.

### parameters


| name | description |
|------|-------------|
| ref | __Required__ The ref to deploy. This can be a branch, tag, or SHA. |
| repo | __Required__ repository in OWNER/REPO form |
| ant-man-preview | The `inactive` state and the `log_url`, `environment_url`, and `auto_inactive` parameters are currently available for developers to preview. Please see the [blog post](https://developer.github.com/changes/2016-04-06-deployment-and-deployment-status-enhancements) for full details.<br><br>To access the API during the preview period, you must set this flag. |
| auto_merge | Attempts to automatically merge the default branch into the requested ref, if it's behind the default branch. |
| created_at |  |
| description | Short description of the deployment. |
| environment | Name for the target deployment environment (e.g., `production`, `staging`, `qa`). |
| payload | JSON payload with extra information about the deployment. |
| production_environment | Specifies if the given environment is one that end-users directly interact with. Default: `true` when `environment` is `production` and `false` otherwise.  <br>**Note:** This parameter requires you to use the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. |
| required_contexts | The [status](https://developer.github.com/v3/repos/statuses/) contexts to verify against commit status checks. If you omit this parameter, GitHub verifies all unique contexts before creating a deployment. To bypass checking entirely, pass an empty array. Defaults to all unique contexts. |
| task | Specifies a task to execute (e.g., `deploy` or `deploy:migrations`). |
| transient_environment | Specifies if the given environment is specific to the deployment and will no longer exist at some point in the future. Default: `false`  <br>**Note:** This parameter requires you to use the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. **Note:** This parameter requires you to use the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. |

## repos create-deployment-status

https://developer.github.com/v3/repos/deployments/#create-a-deployment-status

Users with `push` access can create deployment statuses for a given deployment.

GitHub Apps require `read & write` access to "Deployments" and `read-only` access to "Repo contents" (for private repos). OAuth Apps require the `repo_deployment` scope.

### parameters


| name | description |
|------|-------------|
| deployment_id | __Required__ deployment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| state | __Required__ The state of the status. Can be one of `error`, `failure`, `inactive`, `in_progress`, `queued` `pending`, or `success`. **Note:** To use the `inactive` state, you must provide the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. To use the `in_progress` and `queued` states, you must provide the [`application/vnd.github.flash-preview+json`](https://developer.github.com/v3/previews/#deployment-statuses) custom media type. When you set a transient deployment to `inactive`, the deployment will be shown as `destroyed` in GitHub. |
| ant-man-preview | The `inactive` state and the `log_url`, `environment_url`, and `auto_inactive` parameters are currently available for developers to preview. Please see the [blog post](https://developer.github.com/changes/2016-04-06-deployment-and-deployment-status-enhancements) for full details.<br><br>To access the API during the preview period, you must set this flag. |
| auto_inactive | Adds a new `inactive` status to all prior non-transient, non-production environment deployments with the same repository and `environment` name as the created status's deployment. An `inactive` status is only added to deployments that had a `success` state. Default: `true`  <br>**Note:** To add an `inactive` status to `production` environments, you must use the [`application/vnd.github.flash-preview+json`](https://developer.github.com/v3/previews/#deployment-statuses) custom media type.  <br>**Note:** This parameter requires you to use the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. |
| description | A short description of the status. The maximum description length is 140 characters. |
| environment | Name for the target deployment environment, which can be changed when setting a deploy status. For example, `production`, `staging`, or `qa`. **Note:** This parameter requires you to use the [`application/vnd.github.flash-preview+json`](https://developer.github.com/v3/previews/#deployment-statuses) custom media type. |
| environment_url | Sets the URL for accessing your environment. Default: `""`  <br>**Note:** This parameter requires you to use the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. **Note:** This parameter requires you to use the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. |
| flash-preview | New features in the Deployments API on GitHub are currently available during a public beta. Please see the [blog post](https://developer.github.com/changes/2018-10-16-deployments-environments-states-and-auto-inactive-updates/) for full details.<br><br>To access the new `environment` parameter, the two new values for the `state` parameter (`in_progress` and `queued`), and use `auto_inactive` on production deployments during the public beta period, you must set this flag. |
| log_url | The full URL of the deployment's output. This parameter replaces `target_url`. We will continue to accept `target_url` to support legacy uses, but we recommend replacing `target_url` with `log_url`. Setting `log_url` will automatically set `target_url` to the same value. Default: `""`  <br>**Note:** This parameter requires you to use the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. **Note:** This parameter requires you to use the [`application/vnd.github.ant-man-preview+json`](https://developer.github.com/v3/previews/#enhanced-deployments) custom media type. |
| target_url | The target URL to associate with this status. This URL should contain output to keep the user updated while the task is running or serve as historical information for what happened in the deployment. **Note:** It's recommended to use the `log_url` parameter, which replaces `target_url`. |

## repos create-dispatch-event

https://developer.github.com/v3/repos/#create-a-repository-dispatch-event

You can use this endpoint to trigger a webhook event called `repository_dispatch` when you want activity that happens outside of GitHub to trigger a GitHub Actions workflow or GitHub App webhook. You must configure your GitHub Actions workflow or GitHub App to run when the `repository_dispatch` event occurs. For an example `repository_dispatch` webhook payload, see "[RepositoryDispatchEvent](https://developer.github.com/webhooks/event-payloads/#repository_dispatch)."

The `client_payload` parameter is available for any extra information that your workflow might need. This parameter is a JSON payload that will be passed on when the webhook event is dispatched. For example, the `client_payload` can include a message that a user would like to send using a GitHub Actions workflow. Or the `client_payload` can be used as a test to debug your workflow. For a test example, see the [input example](https://developer.github.com/v3/repos/#example-4).

To give you write access to the repository, you must use a personal access token with the `repo` scope. For more information, see "[Creating a personal access token for the command line](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line)" in the GitHub Help documentation.

This input example shows how you can use the `client_payload` as a test to debug your workflow.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| client_payload | JSON payload with extra information about the webhook event that your action or worklow may use. |
| event_type | **Required:** A custom webhook event name. |

## repos create-for-authenticated-user

https://developer.github.com/v3/repos/#create-a-repository-for-the-authenticated-user

Creates a new repository for the authenticated user.

**OAuth scope requirements**

When using [OAuth](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/), authorizations must include:

*   `public_repo` scope or `repo` scope to create a public repository
*   `repo` scope to create a private repository

### parameters


| name | description |
|------|-------------|
| name | __Required__ The name of the repository. |
| allow_merge_commit | Whether to allow merge commits for pull requests. |
| allow_rebase_merge | Whether to allow rebase merges for pull requests. |
| allow_squash_merge | Whether to allow squash merges for pull requests. |
| auto_init | Whether the repository is initialized with a minimal README. |
| baptiste-preview | The `is_template` and `template_repository` keys are currently available for developer to preview. See [Create a repository using a template](https://developer.github.com/v3/repos/#create-a-repository-using-a-template) to learn how to create template repositories. To access these new response keys during the preview period, you must set this flag. |
| delete_branch_on_merge | Whether to delete head branches when pull requests are merged |
| description | A short description of the repository. |
| gitignore_template | The desired language or platform to apply to the .gitignore. |
| has_downloads | Whether downloads are enabled. |
| has_issues | Whether issues are enabled. |
| has_projects | Whether projects are enabled. |
| has_wiki | Whether the wiki is enabled. |
| homepage | A URL with more information about the repository. |
| is_template | Whether this repository acts as a template that can be used to generate new repositories. |
| license_template | The license keyword of the open source license for this repository. |
| nebula-preview | You can set the visibility of a repository using the new `visibility` parameter in the [Repositories API](https://developer.github.com/v3/repos/), and get a repository's visibility with a new response key. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes/).<br><br>To access repository visibility during the preview period, you must set this flag. |
| private | Whether the repository is private or public. |
| team_id | The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization. |

## repos create-fork

https://developer.github.com/v3/repos/forks/#create-a-fork

Create a fork for the authenticated user.

**Note**: Forking a Repository happens asynchronously. You may have to wait a short period of time before you can access the git objects. If this takes longer than 5 minutes, be sure to contact [GitHub Support](https://github.com/contact) or [GitHub Premium Support](https://premium.githubsupport.com).

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| organization | Optional parameter to specify the organization name if forking into an organization. |

## repos create-in-org

https://developer.github.com/v3/repos/#create-an-organization-repository

Creates a new repository in the specified organization. The authenticated user must be a member of the organization.

**OAuth scope requirements**

When using [OAuth](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/), authorizations must include:

*   `public_repo` scope or `repo` scope to create a public repository
*   `repo` scope to create a private repository

### parameters


| name | description |
|------|-------------|
| name | __Required__ The name of the repository. |
| org | __Required__  |
| allow_merge_commit | Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits. |
| allow_rebase_merge | Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging. |
| allow_squash_merge | Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging. |
| auto_init | Pass `true` to create an initial commit with empty README. |
| baptiste-preview | The `is_template` and `template_repository` keys are currently available for developer to preview. See [Create a repository using a template](https://developer.github.com/v3/repos/#create-a-repository-using-a-template) to learn how to create template repositories. To access these new response keys during the preview period, you must set this flag. |
| delete_branch_on_merge | Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion. |
| description | A short description of the repository. |
| gitignore_template | Desired language or platform [.gitignore template](https://github.com/github/gitignore) to apply. Use the name of the template without the extension. For example, "Haskell". |
| has_issues | Either `true` to enable issues for this repository or `false` to disable them. |
| has_projects | Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error. |
| has_wiki | Either `true` to enable the wiki for this repository or `false` to disable it. |
| homepage | A URL with more information about the repository. |
| is_template | Either `true` to make this repo available as a template repository or `false` to prevent it. |
| license_template | Choose an [open source license template](https://choosealicense.com/) that best suits your needs, and then use the [license keyword](https://help.github.com/articles/licensing-a-repository/#searching-github-by-license-type) as the `license_template` string. For example, "mit" or "mpl-2.0". |
| nebula-preview | You can set the visibility of a repository using the new `visibility` parameter in the [Repositories API](https://developer.github.com/v3/repos/), and get a repository's visibility with a new response key. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes/).<br><br>To access repository visibility during the preview period, you must set this flag. |
| private | Either `true` to create a private repository or `false` to create a public one. |
| team_id | The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization. |
| visibility | Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, `visibility` can also be `internal`. For more information, see "[Creating an internal repository](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/about-repository-visibility#about-internal-repositories)" in the GitHub Help documentation.  <br>The `visibility` parameter overrides the `private` parameter when you use both parameters with the `nebula-preview` preview header. |

## repos create-or-update-file-contents

https://developer.github.com/v3/repos/contents/#create-or-update-file-contents

Creates a new file or replaces an existing file in a repository.

### parameters


| name | description |
|------|-------------|
| content | __Required__ The new file content, using Base64 encoding. |
| message | __Required__ The commit message. |
| path | __Required__ path+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| author.date |  |
| author.email | The email of the author or committer of the commit. You'll receive a `422` status code if `email` is omitted. |
| author.name | The name of the author or committer of the commit. You'll receive a `422` status code if `name` is omitted. |
| branch | The branch name. Default: the repository’s default branch (usually `master`) |
| committer.date |  |
| committer.email | The email of the author or committer of the commit. You'll receive a `422` status code if `email` is omitted. |
| committer.name | The name of the author or committer of the commit. You'll receive a `422` status code if `name` is omitted. |
| sha | **Required if you are updating a file**. The blob SHA of the file being replaced. |

## repos create-pages-site

https://developer.github.com/v3/repos/pages/#create-a-github-pages-site



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| switcheroo-preview | __Required__ Enabling and disabling Pages in the Pages API is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2019-03-14-enabling-disabling-pages/) preview for more details. To access the new endpoints during the preview period, you must set this flag. |
| source.branch | The repository branch used to publish your [site's source files](https://help.github.com/articles/configuring-a-publishing-source-for-github-pages/). Can be either `master` or `gh-pages`. |
| source.path | The repository directory that includes the source files for the Pages site. When `branch` is `master`, you can change `path` to `/docs`. When `branch` is `gh-pages`, you are unable to specify a `path` other than `/`. |

## repos create-release

https://developer.github.com/v3/repos/releases/#create-a-release

Users with push access to the repository can create a release.

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| tag_name | __Required__ The name of the tag. |
| body | Text describing the contents of the tag. |
| draft | `true` to create a draft (unpublished) release, `false` to create a published one. |
| name | The name of the release. |
| prerelease | `true` to identify the release as a prerelease. `false` to identify the release as a full release. |
| target_commitish | Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch (usually `master`). |

## repos create-using-template

https://developer.github.com/v3/repos/#create-a-repository-using-a-template

Creates a new repository using a repository template. Use the `template_owner` and `template_repo` route parameters to specify the repository to use as the template. The authenticated user must own or be a member of an organization that owns the repository. To check if a repository is available to use as a template, get the repository's information using the [Get a repository](https://developer.github.com/v3/repos/#get-a-repository) endpoint and check that the `is_template` key is `true`.

**OAuth scope requirements**

When using [OAuth](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/), authorizations must include:

*   `public_repo` scope or `repo` scope to create a public repository
*   `repo` scope to create a private repository

### parameters


| name | description |
|------|-------------|
| baptiste-preview | __Required__ The `is_template` and `template_repository` keys are currently available for developer to preview. See [Create a repository using a template](https://developer.github.com/v3/repos/#create-a-repository-using-a-template) to learn how to create template repositories. To access these new response keys during the preview period, you must set this flag. |
| name | __Required__ The name of the new repository. |
| template_owner | __Required__ template_owner parameter |
| template_repo | __Required__ template_repo parameter |
| description | A short description of the new repository. |
| owner | The organization or person who will own the new repository. To create a new repository in an organization, the authenticated user must be a member of the specified organization. |
| private | Either `true` to create a new private repository or `false` to create a new public one. |

## repos create-webhook

https://developer.github.com/v3/repos/hooks/#create-a-repository-webhook

Repositories can have multiple webhooks installed. Each webhook should have a unique `config`. Multiple webhooks can
share the same `config` as long as those webhooks do not have any `events` that overlap.

### parameters


| name | description |
|------|-------------|
| config.url | __Required__ The URL to which the payloads will be delivered. |
| repo | __Required__ repository in OWNER/REPO form |
| active | Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications. |
| config.content_type | The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`. |
| config.digest |  |
| config.insecure_ssl | Determines whether the SSL certificate of the host for `url` will be verified when delivering payloads. Supported values include `0` (verification is performed) and `1` (verification is not performed). The default is `0`. **We strongly recommend not setting this to `1` as you are subject to man-in-the-middle and other attacks.** |
| config.secret | If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value in the [`X-Hub-Signature`](https://developer.github.com/webhooks/event-payloads/#delivery-headers) header. |
| config.token |  |
| events | Determines what [events](https://developer.github.com/webhooks/event-payloads) the hook is triggered for. |
| name | Use `web` to create a webhook. Default: `web`. This parameter only accepts the value `web`. |

## repos decline-invitation

https://developer.github.com/v3/repos/invitations/#decline-a-repository-invitation



### parameters


| name | description |
|------|-------------|
| invitation_id | __Required__ invitation_id parameter |

## repos delete

https://developer.github.com/v3/repos/#delete-a-repository

Deleting a repository requires admin access. If OAuth is used, the `delete_repo` scope is required.

If an organization owner has configured the organization to prevent members from deleting organization-owned
repositories, you will get a `403 Forbidden` response.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-access-restrictions

https://developer.github.com/v3/repos/branches/#delete-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Disables the ability to restrict who can push to this branch.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-admin-branch-protection

https://developer.github.com/v3/repos/branches/#delete-admin-branch-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Removing admin enforcement requires admin or owner permissions to the repository and branch protection to be enabled.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-branch-protection

https://developer.github.com/v3/repos/branches/#delete-branch-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-commit-comment

https://developer.github.com/v3/repos/comments/#delete-a-commit-comment



### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-commit-signature-protection

https://developer.github.com/v3/repos/branches/#delete-commit-signature-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

When authenticated with admin or owner permissions to the repository, you can use this endpoint to disable required signed commits on a branch. You must enable branch protection to require signed commits.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| zzzax-preview | __Required__ Protected Branches API can now manage a setting for requiring signed commits. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-02-22-protected-branches-required-signatures) for full details. To access the API during the preview period, you must set this flag. |

## repos delete-deploy-key

https://developer.github.com/v3/repos/keys/#delete-a-deploy-key

Deploy keys are immutable. If you need to update a key, remove the key and create a new one instead.

### parameters


| name | description |
|------|-------------|
| key_id | __Required__ key_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-deployment

https://developer.github.com/v3/repos/deployments/#delete-a-deployment

To ensure there can always be an active deployment, you can only delete an _inactive_ deployment. Anyone with `repo` or `repo_deployment` scopes can delete an inactive deployment.

To set a deployment as inactive, you must:

*   Create a new deployment that is active so that the system has a record of the current state, then delete the previously active deployment.
*   Mark the active deployment as inactive by adding any non-successful deployment status.

For more information, see "[Create a deployment](https://developer.github.com/v3/repos/deployments/#create-a-deployment)" and "[Create a deployment status](https://developer.github.com/v3/repos/deployments/#create-a-deployment-status)."

### parameters


| name | description |
|------|-------------|
| deployment_id | __Required__ deployment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-file

https://developer.github.com/v3/repos/contents/#delete-a-file

Deletes a file in a repository.

You can provide an additional `committer` parameter, which is an object containing information about the committer. Or, you can provide an `author` parameter, which is an object containing information about the author.

The `author` section is optional and is filled in with the `committer` information if omitted. If the `committer` information is omitted, the authenticated user's information is used.

You must provide values for both `name` and `email`, whether you choose to use `author` or `committer`. Otherwise, you'll receive a `422` status code.

### parameters


| name | description |
|------|-------------|
| message | __Required__ The commit message. |
| path | __Required__ path+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| sha | __Required__ The blob SHA of the file being replaced. |
| author.email | The email of the author (or committer) of the commit |
| author.name | The name of the author (or committer) of the commit |
| branch | The branch name. Default: the repository’s default branch (usually `master`) |
| committer.email | The email of the author (or committer) of the commit |
| committer.name | The name of the author (or committer) of the commit |

## repos delete-invitation

https://developer.github.com/v3/repos/invitations/#delete-a-repository-invitation



### parameters


| name | description |
|------|-------------|
| invitation_id | __Required__ invitation_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-pages-site

https://developer.github.com/v3/repos/pages/#delete-a-github-pages-site



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| switcheroo-preview | __Required__ Enabling and disabling Pages in the Pages API is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2019-03-14-enabling-disabling-pages/) preview for more details. To access the new endpoints during the preview period, you must set this flag. |

## repos delete-pull-request-review-protection

https://developer.github.com/v3/repos/branches/#delete-pull-request-review-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-release

https://developer.github.com/v3/repos/releases/#delete-a-release

Users with push access to the repository can delete a release.

### parameters


| name | description |
|------|-------------|
| release_id | __Required__ release_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-release-asset

https://developer.github.com/v3/repos/releases/#delete-a-release-asset



### parameters


| name | description |
|------|-------------|
| asset_id | __Required__ asset_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos delete-webhook

https://developer.github.com/v3/repos/hooks/#delete-a-repository-webhook



### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |

## repos disable-automated-security-fixes

https://developer.github.com/v3/repos/#disable-automated-security-fixes

Disables automated security fixes for a repository. The authenticated user must have admin access to the repository. For more information, see "[Configuring automated security fixes](https://help.github.com/en/articles/configuring-automated-security-fixes)".

### parameters


| name | description |
|------|-------------|
| london-preview | __Required__ Enabling or disabling automated security fixes is currently available for developers to preview. To access this new endpoint during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |

## repos disable-vulnerability-alerts

https://developer.github.com/v3/repos/#disable-vulnerability-alerts

Disables dependency alerts and the dependency graph for a repository. The authenticated user must have admin access to the repository. For more information, see "[About security alerts for vulnerable dependencies](https://help.github.com/en/articles/about-security-alerts-for-vulnerable-dependencies)".

### parameters


| name | description |
|------|-------------|
| dorian-preview | __Required__ Enabling and disabling dependency alerts for a repository using the REST API is currently available for developers to preview. To access these new endpoints during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |

## repos download-tarball-archive

https://developer.github.com/v3/repos/contents/#download-a-repository-archive

Gets a redirect URL to download a tar archive for a repository. If you omit `:ref`, the repository’s default branch (usually
`master`) will be used. Please make sure your HTTP framework is configured to follow redirects or you will need to use
the `Location` header to make a second `GET` request.
**Note**: For private repositories, these links are temporary and expire after five minutes.

### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos download-zipball-archive

https://developer.github.com/v3/repos/contents/#download-a-repository-archive

Gets a redirect URL to download a zip archive for a repository. If you omit `:ref`, the repository’s default branch (usually
`master`) will be used. Please make sure your HTTP framework is configured to follow redirects or you will need to use
the `Location` header to make a second `GET` request.
**Note**: For private repositories, these links are temporary and expire after five minutes.

### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos enable-automated-security-fixes

https://developer.github.com/v3/repos/#enable-automated-security-fixes

Enables automated security fixes for a repository. The authenticated user must have admin access to the repository. For more information, see "[Configuring automated security fixes](https://help.github.com/en/articles/configuring-automated-security-fixes)".

### parameters


| name | description |
|------|-------------|
| london-preview | __Required__ Enabling or disabling automated security fixes is currently available for developers to preview. To access this new endpoint during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |

## repos enable-vulnerability-alerts

https://developer.github.com/v3/repos/#enable-vulnerability-alerts

Enables dependency alerts and the dependency graph for a repository. The authenticated user must have admin access to the repository. For more information, see "[About security alerts for vulnerable dependencies](https://help.github.com/en/articles/about-security-alerts-for-vulnerable-dependencies)".

### parameters


| name | description |
|------|-------------|
| dorian-preview | __Required__ Enabling and disabling dependency alerts for a repository using the REST API is currently available for developers to preview. To access these new endpoints during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |

## repos get

https://developer.github.com/v3/repos/#get-a-repository

When you pass the `scarlet-witch-preview` media type, requests to get a repository will also return the repository's code of conduct if it can be detected from the repository's code of conduct file.

The `parent` and `source` objects are present when the repository is a fork. `parent` is the repository this repository was forked from, `source` is the ultimate source for the network.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| nebula-preview | You can set the visibility of a repository using the new `visibility` parameter in the [Repositories API](https://developer.github.com/v3/repos/), and get a repository's visibility with a new response key. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes/).<br><br>To access repository visibility during the preview period, you must set this flag. |
| scarlet-witch-preview | The Codes of Conduct API is currently available for developers to preview.<br><br>To access the API during the preview period, you must set this flag. |

## repos get-access-restrictions

https://developer.github.com/v3/repos/branches/#get-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Lists who has access to this protected branch.

**Note**: Users, apps, and teams `restrictions` are only available for organization-owned repositories.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-admin-branch-protection

https://developer.github.com/v3/repos/branches/#get-admin-branch-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-all-status-check-contexts

https://developer.github.com/v3/repos/branches/#get-all-status-check-contexts

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-all-topics

https://developer.github.com/v3/repos/#get-all-repository-topics



### parameters


| name | description |
|------|-------------|
| mercy-preview | __Required__ The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-apps-with-access-to-protected-branch

https://developer.github.com/v3/repos/branches/#list-apps-with-access-to-the-protected-branch

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Lists the GitHub Apps that have push access to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-branch

https://developer.github.com/v3/repos/branches/#get-a-branch



### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-branch-protection

https://developer.github.com/v3/repos/branches/#get-branch-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| luke-cage-preview | The Protected Branches API now has a setting for requiring a specified number of approving pull request reviews before merging. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-03-16-protected-branches-required-approving-reviews) for full details. To access the API during the preview period, you must set this flag. |

## repos get-clones

https://developer.github.com/v3/repos/traffic/#get-repository-clones

Get the total number of clones and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| per | Must be one of: `day`, `week`. |

## repos get-code-frequency-stats

https://developer.github.com/v3/repos/statistics/#get-the-weekly-commit-activity

Returns a weekly aggregate of the number of additions and deletions pushed to a repository.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-collaborator-permission-level

https://developer.github.com/v3/repos/collaborators/#get-repository-permissions-for-a-user

Checks the repository permission of a collaborator. The possible repository permissions are `admin`, `write`, `read`, and `none`.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| username | __Required__  |

## repos get-combined-status-for-ref

https://developer.github.com/v3/repos/statuses/#get-the-combined-status-for-a-specific-reference

Users with pull access in a repository can access a combined view of commit statuses for a given ref. The ref can be a SHA, a branch name, or a tag name.

The most recent status for each context is returned, up to 100. This field [paginates](https://developer.github.com/v3/#pagination) if there are over 100 contexts.

Additionally, a combined `state` is returned. The `state` is one of:

*   **failure** if any of the contexts report as `error` or `failure`
*   **pending** if there are no statuses or a context is `pending`
*   **success** if the latest status for all contexts is `success`

### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-commit

https://developer.github.com/v3/repos/commits/#get-a-commit

Returns the contents of a single commit reference. You must have `read` access for the repository to use this endpoint.

You can pass the appropriate [media type](https://developer.github.com/v3/media/#commits-commit-comparison-and-pull-requests) to fetch `diff` and `patch` formats. Diffs with binary data will have no `patch` property.

To return only the SHA-1 hash of the commit reference, you can provide the `sha` custom [media type](https://developer.github.com/v3/media/#commits-commit-comparison-and-pull-requests) in the `Accept` header. You can use this endpoint to check if a remote reference's SHA-1 hash is the same as your local reference's SHA-1 hash by providing the local SHA-1 reference as the ETag.

**Signature verification object**

The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:

These are the possible values for `reason` in the `verification` object:

| Value                    | Description                                                                                                                       |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| `expired_key`            | The key that made the signature is expired.                                                                                       |
| `not_signing_key`        | The "signing" flag is not among the usage flags in the GPG key that made the signature.                                           |
| `gpgverify_error`        | There was an error communicating with the signature verification service.                                                         |
| `gpgverify_unavailable`  | The signature verification service is currently unavailable.                                                                      |
| `unsigned`               | The object does not include a signature.                                                                                          |
| `unknown_signature_type` | A non-PGP signature was found in the commit.                                                                                      |
| `no_user`                | No user was associated with the `committer` email address in the commit.                                                          |
| `unverified_email`       | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. |
| `bad_email`              | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature.             |
| `unknown_key`            | The key that made the signature has not been registered with any user's account.                                                  |
| `malformed_signature`    | There was an error parsing the signature.                                                                                         |
| `invalid`                | The signature could not be cryptographically verified using the key whose key-id was found in the signature.                      |
| `valid`                  | None of the above errors applied, so the signature is considered to be verified.                                                  |

### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-commit-activity-stats

https://developer.github.com/v3/repos/statistics/#get-the-last-year-of-commit-activity

Returns the last year of commit activity grouped by week. The `days` array is a group of commits per day, starting on `Sunday`.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-commit-comment

https://developer.github.com/v3/repos/comments/#get-a-commit-comment



### parameters


| name | description |
|------|-------------|
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## repos get-commit-signature-protection

https://developer.github.com/v3/repos/branches/#get-commit-signature-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

When authenticated with admin or owner permissions to the repository, you can use this endpoint to check whether a branch requires signed commits. An enabled status of `true` indicates you must sign commits on this branch. For more information, see [Signing commits with GPG](https://help.github.com/articles/signing-commits-with-gpg) in GitHub Help.

**Note**: You must enable branch protection to require signed commits.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| zzzax-preview | __Required__ Protected Branches API can now manage a setting for requiring signed commits. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-02-22-protected-branches-required-signatures) for full details. To access the API during the preview period, you must set this flag. |

## repos get-community-profile-metrics

https://developer.github.com/v3/repos/community/#get-community-profile-metrics

This endpoint will return all community profile metrics, including an overall health score, repository description, the presence of documentation, detected code of conduct, detected license, and the presence of ISSUE\_TEMPLATE, PULL\_REQUEST\_TEMPLATE, README, and CONTRIBUTING files.

### parameters


| name | description |
|------|-------------|
| black-panther-preview | __Required__ We're currently offering a preview of the Community Profile API (also known as community health). To access the API during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-content

https://developer.github.com/v3/repos/contents/#get-repository-content

Gets the contents of a file or directory in a repository. Specify the file path or directory in `:path`. If you omit
`:path`, you will receive the contents of all files in the repository.

Files and symlinks support [a custom media type](https://developer.github.com/v3/repos/contents/#custom-media-types) for
retrieving the raw content or rendered HTML (when supported). All content types support [a custom media
type](https://developer.github.com/v3/repos/contents/#custom-media-types) to ensure the content is returned in a consistent
object format.

**Note**:
*   To get a repository's contents recursively, you can [recursively get the tree](https://developer.github.com/v3/git/trees/).
*   This API has an upper limit of 1,000 files for a directory. If you need to retrieve more files, use the [Git Trees
API](https://developer.github.com/v3/git/trees/#get-a-tree).
*   This API supports files up to 1 megabyte in size.

#### If the content is a directory
The response will be an array of objects, one object for each item in the directory.
When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value
_should_ be "submodule". This behavior exists in API v3 [for backwards compatibility purposes](https://git.io/v1YCW).
In the next major version of the API, the type will be returned as "submodule".

#### If the content is a symlink 
If the requested `:path` points to a symlink, and the symlink's target is a normal file in the repository, then the
API responds with the content of the file (in the format shown in the example. Otherwise, the API responds with an object 
describing the symlink itself.

#### If the content is a submodule
The `submodule_git_url` identifies the location of the submodule repository, and the `sha` identifies a specific
commit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks out
the submodule at that specific commit.

If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and the
github.com URLs (`html_url` and `_links["html"]`) will have null values.

### parameters


| name | description |
|------|-------------|
| path | __Required__ path+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| ref | The name of the commit/branch/tag. Default: the repository’s default branch (usually `master`) |

## repos get-contributors-stats

https://developer.github.com/v3/repos/statistics/#get-all-contributor-commit-activity


Returns the `total` number of commits authored by the contributor. In addition, the response includes a Weekly Hash (`weeks` array) with the following information:

*   `w` - Start of the week, given as a [Unix timestamp](http://en.wikipedia.org/wiki/Unix_time).
*   `a` - Number of additions
*   `d` - Number of deletions
*   `c` - Number of commits

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-deploy-key

https://developer.github.com/v3/repos/keys/#get-a-deploy-key



### parameters


| name | description |
|------|-------------|
| key_id | __Required__ key_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-deployment

https://developer.github.com/v3/repos/deployments/#get-a-deployment



### parameters


| name | description |
|------|-------------|
| deployment_id | __Required__ deployment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| ant-man-preview | The `inactive` state and the `log_url`, `environment_url`, and `auto_inactive` parameters are currently available for developers to preview. Please see the [blog post](https://developer.github.com/changes/2016-04-06-deployment-and-deployment-status-enhancements) for full details.<br><br>To access the API during the preview period, you must set this flag. |
| machine-man-preview | To access the API with your GitHub App, you must set this flag for your requests. |

## repos get-deployment-status

https://developer.github.com/v3/repos/deployments/#get-a-deployment-status

Users with pull access can view a deployment status for a deployment:

### parameters


| name | description |
|------|-------------|
| deployment_id | __Required__ deployment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| status_id | __Required__ status_id parameter |
| ant-man-preview | The `inactive` state and the `log_url`, `environment_url`, and `auto_inactive` parameters are currently available for developers to preview. Please see the [blog post](https://developer.github.com/changes/2016-04-06-deployment-and-deployment-status-enhancements) for full details.<br><br>To access the API during the preview period, you must set this flag. |
| flash-preview | New features in the Deployments API on GitHub are currently available during a public beta. Please see the [blog post](https://developer.github.com/changes/2018-10-16-deployments-environments-states-and-auto-inactive-updates/) for full details.<br><br>To access the new `environment` parameter, the two new values for the `state` parameter (`in_progress` and `queued`), and use `auto_inactive` on production deployments during the public beta period, you must set this flag. |
| machine-man-preview | To access the API with your GitHub App, you must set this flag for your requests. |

## repos get-latest-pages-build

https://developer.github.com/v3/repos/pages/#get-latest-pages-build



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-latest-release

https://developer.github.com/v3/repos/releases/#get-the-latest-release

View the latest published full release for the repository.

The latest release is the most recent non-prerelease, non-draft release, sorted by the `created_at` attribute. The `created_at` attribute is the date of the commit used for the release, and not the date when the release was drafted or published.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-pages

https://developer.github.com/v3/repos/pages/#get-a-github-pages-site



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-pages-build

https://developer.github.com/v3/repos/pages/#get-github-pages-build



### parameters


| name | description |
|------|-------------|
| build_id | __Required__ build_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-participation-stats

https://developer.github.com/v3/repos/statistics/#get-the-weekly-commit-count

Returns the total commit counts for the `owner` and total commit counts in `all`. `all` is everyone combined, including the `owner` in the last 52 weeks. If you'd like to get the commit counts for non-owners, you can subtract `owner` from `all`.

The array order is oldest week (index 0) to most recent week.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-pull-request-review-protection

https://developer.github.com/v3/repos/branches/#get-pull-request-review-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| luke-cage-preview | The Protected Branches API now has a setting for requiring a specified number of approving pull request reviews before merging. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-03-16-protected-branches-required-approving-reviews) for full details. To access the API during the preview period, you must set this flag. |

## repos get-punch-card-stats

https://developer.github.com/v3/repos/statistics/#get-the-hourly-commit-count-for-each-day

Each array contains the day number, hour number, and number of commits:

*   `0-6`: Sunday - Saturday
*   `0-23`: Hour of day
*   Number of commits

For example, `[2, 14, 25]` indicates that there were 25 total commits, during the 2:00pm hour on Tuesdays. All times are based on the time zone of individual commits.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-readme

https://developer.github.com/v3/repos/contents/#get-a-repository-readme

Gets the preferred README for a repository.

READMEs support [custom media types](https://developer.github.com/v3/repos/contents/#custom-media-types) for retrieving the raw content or rendered HTML.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| ref | The name of the commit/branch/tag. Default: the repository’s default branch (usually `master`) |

## repos get-release

https://developer.github.com/v3/repos/releases/#get-a-release

**Note:** This returns an `upload_url` key corresponding to the endpoint for uploading release assets. This key is a [hypermedia resource](https://developer.github.com/v3/#hypermedia).

### parameters


| name | description |
|------|-------------|
| release_id | __Required__ release_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-release-asset

https://developer.github.com/v3/repos/releases/#get-a-release-asset

To download the asset's binary content, set the `Accept` header of the request to [`application/octet-stream`](https://developer.github.com/v3/media/#media-types). The API will either redirect the client to the location, or stream it directly if possible. API clients should handle both a `200` or `302` response.

### parameters


| name | description |
|------|-------------|
| asset_id | __Required__ asset_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-release-by-tag

https://developer.github.com/v3/repos/releases/#get-a-release-by-tag-name

Get a published release with the specified tag.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| tag | __Required__ tag+ parameter |

## repos get-status-checks-protection

https://developer.github.com/v3/repos/branches/#get-status-checks-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-teams-with-access-to-protected-branch

https://developer.github.com/v3/repos/branches/#list-teams-with-access-to-the-protected-branch

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Lists the teams who have push access to this branch. The list includes child teams.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-top-paths

https://developer.github.com/v3/repos/traffic/#get-top-referral-paths

Get the top 10 popular contents over the last 14 days.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-top-referrers

https://developer.github.com/v3/repos/traffic/#get-top-referral-sources

Get the top 10 referrers over the last 14 days.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos get-users-with-access-to-protected-branch

https://developer.github.com/v3/repos/branches/#list-users-with-access-to-the-protected-branch

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Lists the people who have push access to this branch.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos get-views

https://developer.github.com/v3/repos/traffic/#get-page-views

Get the total number of views and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| per | Must be one of: `day`, `week`. |

## repos get-webhook

https://developer.github.com/v3/repos/hooks/#get-a-repository-webhook



### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |

## repos list-branches

https://developer.github.com/v3/repos/branches/#list-branches



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| protected | Setting to `true` returns only protected branches. When set to `false`, only unprotected branches are returned. Omitting this parameter returns all branches. |

## repos list-branches-for-head-commit

https://developer.github.com/v3/repos/commits/#list-branches-for-head-commit

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Returns all branches where the given commit SHA is the HEAD, or latest commit for the branch.

### parameters


| name | description |
|------|-------------|
| commit_sha | __Required__ commit_sha+ parameter |
| groot-preview | __Required__ Listing branches or pull requests for a commit in the Commits API is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2019-04-11-pulls-branches-for-commit/) for more details. To access the new endpoints during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |

## repos list-collaborators

https://developer.github.com/v3/repos/collaborators/#list-repository-collaborators

For organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.

Team members will include the members of child teams.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| affiliation | Filter collaborators returned by their affiliation. Can be one of:  <br>\* `outside`: All outside collaborators of an organization-owned repository.  <br>\* `direct`: All collaborators with permissions to an organization-owned repository, regardless of organization membership status.  <br>\* `all`: All collaborators the authenticated user can see. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-comments-for-commit

https://developer.github.com/v3/repos/comments/#list-commit-comments

Use the `:commit_sha` to specify the commit that will have its comments listed.

### parameters


| name | description |
|------|-------------|
| commit_sha | __Required__ commit_sha+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## repos list-commit-comments-for-repo

https://developer.github.com/v3/repos/comments/#list-commit-comments-for-a-repository

Commit Comments use [these custom media types](https://developer.github.com/v3/repos/comments/#custom-media-types). You can read more about the use of media types in the API [here](https://developer.github.com/v3/media/).

Comments are ordered by ascending ID.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## repos list-commit-statuses-for-ref

https://developer.github.com/v3/repos/statuses/#list-commit-statuses-for-a-reference

Users with pull access in a repository can view commit statuses for a given ref. The ref can be a SHA, a branch name, or a tag name. Statuses are returned in reverse chronological order. The first status in the list will be the latest one.

This resource is also available via a legacy route: `GET /repos/:owner/:repo/statuses/:ref`.

### parameters


| name | description |
|------|-------------|
| ref | __Required__ ref+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-commits

https://developer.github.com/v3/repos/commits/#list-commits

**Signature verification object**

The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:

These are the possible values for `reason` in the `verification` object:

| Value                    | Description                                                                                                                       |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| `expired_key`            | The key that made the signature is expired.                                                                                       |
| `not_signing_key`        | The "signing" flag is not among the usage flags in the GPG key that made the signature.                                           |
| `gpgverify_error`        | There was an error communicating with the signature verification service.                                                         |
| `gpgverify_unavailable`  | The signature verification service is currently unavailable.                                                                      |
| `unsigned`               | The object does not include a signature.                                                                                          |
| `unknown_signature_type` | A non-PGP signature was found in the commit.                                                                                      |
| `no_user`                | No user was associated with the `committer` email address in the commit.                                                          |
| `unverified_email`       | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. |
| `bad_email`              | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature.             |
| `unknown_key`            | The key that made the signature has not been registered with any user's account.                                                  |
| `malformed_signature`    | There was an error parsing the signature.                                                                                         |
| `invalid`                | The signature could not be cryptographically verified using the key whose key-id was found in the signature.                      |
| `valid`                  | None of the above errors applied, so the signature is considered to be verified.                                                  |

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| author | GitHub login or email address by which to filter by commit author. |
| page | Page number of the results to fetch. |
| path | Only commits containing this file path will be returned. |
| per_page | Results per page (max 100) |
| sha | SHA or branch to start listing commits from. Default: the repository’s default branch (usually `master`). |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| until | Only commits before this date will be returned. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## repos list-contributors

https://developer.github.com/v3/repos/#list-repository-contributors

Lists contributors to the specified repository and sorts them by the number of commits per contributor in descending order. This endpoint may return information that is a few hours old because the GitHub REST API v3 caches contributor data to improve performance.

GitHub identifies contributors by author email address. This endpoint groups contribution counts by GitHub user, which includes all associated email addresses. To improve performance, only the first 500 author email addresses in the repository link to GitHub users. The rest will appear as anonymous contributors without associated GitHub user information.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| anon | Set to `1` or `true` to include anonymous contributors in results. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-deploy-keys

https://developer.github.com/v3/repos/keys/#list-deploy-keys



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-deployment-statuses

https://developer.github.com/v3/repos/deployments/#list-deployment-statuses

Users with pull access can view deployment statuses for a deployment:

### parameters


| name | description |
|------|-------------|
| deployment_id | __Required__ deployment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| ant-man-preview | The `inactive` state and the `log_url`, `environment_url`, and `auto_inactive` parameters are currently available for developers to preview. Please see the [blog post](https://developer.github.com/changes/2016-04-06-deployment-and-deployment-status-enhancements) for full details.<br><br>To access the API during the preview period, you must set this flag. |
| flash-preview | New features in the Deployments API on GitHub are currently available during a public beta. Please see the [blog post](https://developer.github.com/changes/2018-10-16-deployments-environments-states-and-auto-inactive-updates/) for full details.<br><br>To access the new `environment` parameter, the two new values for the `state` parameter (`in_progress` and `queued`), and use `auto_inactive` on production deployments during the public beta period, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-deployments

https://developer.github.com/v3/repos/deployments/#list-deployments

Simple filtering of deployments is available via query parameters:

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| ant-man-preview | The `inactive` state and the `log_url`, `environment_url`, and `auto_inactive` parameters are currently available for developers to preview. Please see the [blog post](https://developer.github.com/changes/2016-04-06-deployment-and-deployment-status-enhancements) for full details.<br><br>To access the API during the preview period, you must set this flag. |
| environment | The name of the environment that was deployed to (e.g., `staging` or `production`). |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| ref | The name of the ref. This can be a branch, tag, or SHA. |
| sha | The SHA recorded at creation time. |
| task | The name of the task for the deployment (e.g., `deploy` or `deploy:migrations`). |

## repos list-for-authenticated-user

https://developer.github.com/v3/repos/#list-repositories-for-the-authenticated-user

Lists repositories that the authenticated user has explicit permission (`:read`, `:write`, or `:admin`) to access.

The authenticated user has explicit permission to access repositories they own, repositories where they are a collaborator, and repositories that they can access through an organization membership.

### parameters


| name | description |
|------|-------------|
| affiliation | Comma-separated list of values. Can include:  <br>\* `owner`: Repositories that are owned by the authenticated user.  <br>\* `collaborator`: Repositories that the user has been added to as a collaborator.  <br>\* `organization_member`: Repositories that the user has access to through being a member of an organization. This includes every repository on every team that the user is on. |
| before | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| direction | Can be one of `asc` or `desc`. Default: `asc` when using `full_name`, otherwise `desc` |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| sort | Can be one of `created`, `updated`, `pushed`, `full_name`. |
| type | Can be one of `all`, `owner`, `public`, `private`, `member`. Default: `all`  <br>  <br>Will cause a `422` error if used in the same request as **visibility** or **affiliation**. Will cause a `422` error if used in the same request as **visibility** or **affiliation**. |
| visibility | Can be one of `all`, `public`, or `private`. |

## repos list-for-org

https://developer.github.com/v3/repos/#list-organization-repositories

Lists repositories for the specified organization.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| baptiste-preview | The `is_template` and `template_repository` keys are currently available for developer to preview. See [Create a repository using a template](https://developer.github.com/v3/repos/#create-a-repository-using-a-template) to learn how to create template repositories. To access these new response keys during the preview period, you must set this flag. |
| direction | Can be one of `asc` or `desc`. Default: when using `full_name`: `asc`, otherwise `desc` |
| nebula-preview | You can set the visibility of a repository using the new `visibility` parameter in the [Repositories API](https://developer.github.com/v3/repos/), and get a repository's visibility with a new response key. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes/).<br><br>To access repository visibility during the preview period, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | Can be one of `created`, `updated`, `pushed`, `full_name`. |
| type | Specifies the types of repositories you want returned. Can be one of `all`, `public`, `private`, `forks`, `sources`, `member`, `internal`. Default: `all`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, `type` can also be `internal`. |

## repos list-for-user

https://developer.github.com/v3/repos/#list-repositories-for-a-user

Lists public repositories for the specified user.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| direction | Can be one of `asc` or `desc`. Default: `asc` when using `full_name`, otherwise `desc` |
| nebula-preview | You can set the visibility of a repository using the new `visibility` parameter in the [Repositories API](https://developer.github.com/v3/repos/), and get a repository's visibility with a new response key. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes/).<br><br>To access repository visibility during the preview period, you must set this flag. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | Can be one of `created`, `updated`, `pushed`, `full_name`. |
| type | Can be one of `all`, `owner`, `member`. |

## repos list-forks

https://developer.github.com/v3/repos/forks/#list-forks



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | The sort order. Can be either `newest`, `oldest`, or `stargazers`. |

## repos list-invitations

https://developer.github.com/v3/repos/invitations/#list-repository-invitations

When authenticating as a user with admin rights to a repository, this endpoint will list all currently open repository invitations.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-invitations-for-authenticated-user

https://developer.github.com/v3/repos/invitations/#list-repository-invitations-for-the-authenticated-user

When authenticating as a user, this endpoint will list all currently open repository invitations for that user.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-languages

https://developer.github.com/v3/repos/#list-repository-languages

Lists languages for the specified repository. The value shown for each language is the number of bytes of code written in that language.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos list-pages-builds

https://developer.github.com/v3/repos/pages/#list-github-pages-builds



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-public

https://developer.github.com/v3/repos/#list-public-repositories

Lists all public repositories in the order that they were created.

Note: Pagination is powered exclusively by the `since` parameter. Use the [Link header](https://developer.github.com/v3/#link-header) to get the URL for the next page of repositories.

### parameters


| name | description |
|------|-------------|
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |
| visibility |  |

## repos list-pull-requests-associated-with-commit

https://developer.github.com/v3/repos/commits/#list-pull-requests-associated-with-a-commit

Lists all pull requests containing the provided commit SHA, which can be from any point in the commit history. The results will include open and closed pull requests. Additional preview headers may be required to see certain details for associated pull requests, such as whether a pull request is in a draft state. For more information about previews that might affect this endpoint, see the [List pull requests](https://developer.github.com/v3/pulls/#list-pull-requests) endpoint.

### parameters


| name | description |
|------|-------------|
| commit_sha | __Required__ commit_sha+ parameter |
| groot-preview | __Required__ Listing branches or pull requests for a commit in the Commits API is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2019-04-11-pulls-branches-for-commit/) for more details. To access the new endpoints during the preview period, you must set this flag. |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-release-assets

https://developer.github.com/v3/repos/releases/#list-release-assets



### parameters


| name | description |
|------|-------------|
| release_id | __Required__ release_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-releases

https://developer.github.com/v3/repos/releases/#list-releases

This returns a list of releases, which does not include regular Git tags that have not been associated with a release. To get a list of Git tags, use the [Repository Tags API](https://developer.github.com/v3/repos/#list-repository-tags).

Information about published releases are available to everyone. Only users with push access will receive listings for draft releases.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-tags

https://developer.github.com/v3/repos/#list-repository-tags



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-teams

https://developer.github.com/v3/repos/#list-repository-teams



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos list-webhooks

https://developer.github.com/v3/repos/hooks/#list-repository-webhooks



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## repos merge

https://developer.github.com/v3/repos/merging/#merge-a-branch



### parameters


| name | description |
|------|-------------|
| base | __Required__ The name of the base branch that the head will be merged into. |
| head | __Required__ The head to merge. This can be a branch name or a commit SHA1. |
| repo | __Required__ repository in OWNER/REPO form |
| commit_message | Commit message to use for the merge commit. If omitted, a default message will be used. |

## repos ping-webhook

https://developer.github.com/v3/repos/hooks/#ping-a-repository-webhook

This will trigger a [ping event](https://developer.github.com/webhooks/#ping-event) to be sent to the hook.

### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |

## repos remove-app-access-restrictions

https://developer.github.com/v3/repos/branches/#remove-app-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Removes the ability of an app to push to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.

| Type    | Description                                                                                                                                                |
| ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `array` | The GitHub Apps that have push access to this branch. Use the app's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos remove-collaborator

https://developer.github.com/v3/repos/collaborators/#remove-a-repository-collaborator



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| username | __Required__  |

## repos remove-status-check-contexts

https://developer.github.com/v3/repos/branches/#remove-status-check-contexts

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos remove-status-check-protection

https://developer.github.com/v3/repos/branches/#remove-status-check-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos remove-team-access-restrictions

https://developer.github.com/v3/repos/branches/#remove-team-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Removes the ability of a team to push to this branch. You can also remove push access for child teams.

| Type    | Description                                                                                                                                         |
| ------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `array` | Teams that should no longer have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos remove-user-access-restrictions

https://developer.github.com/v3/repos/branches/#remove-user-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Removes the ability of a user to push to this branch.

| Type    | Description                                                                                                                                   |
| ------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `array` | Usernames of the people who should no longer have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos replace-all-topics

https://developer.github.com/v3/repos/#replace-all-repository-topics



### parameters


| name | description |
|------|-------------|
| mercy-preview | __Required__ The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must set this flag. |
| names | __Required__ An array of topics to add to the repository. Pass one or more topics to _replace_ the set of existing topics. Send an empty array (`[]`) to clear all topics from the repository. **Note:** Topic `names` cannot contain uppercase letters. |
| repo | __Required__ repository in OWNER/REPO form |

## repos request-pages-build

https://developer.github.com/v3/repos/pages/#request-a-github-pages-build

You can request that your site be built from the latest revision on the default branch. This has the same effect as pushing a commit to your default branch, but does not require an additional commit. Manually triggering page builds can be helpful when diagnosing build warnings and failures.

Build requests are limited to one concurrent build per repository and one concurrent build per requester. If you request a build while another is still in progress, the second request will be queued until the first completes.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |

## repos set-admin-branch-protection

https://developer.github.com/v3/repos/branches/#set-admin-branch-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Adding admin enforcement requires admin or owner permissions to the repository and branch protection to be enabled.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos set-app-access-restrictions

https://developer.github.com/v3/repos/branches/#set-app-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Replaces the list of apps that have push access to this branch. This removes all apps that previously had push access and grants push access to the new list of apps. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.

| Type    | Description                                                                                                                                                |
| ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `array` | The GitHub Apps that have push access to this branch. Use the app's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos set-status-check-contexts

https://developer.github.com/v3/repos/branches/#set-status-check-contexts

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos set-team-access-restrictions

https://developer.github.com/v3/repos/branches/#set-team-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Replaces the list of teams that have push access to this branch. This removes all teams that previously had push access and grants push access to the new list of teams. Team restrictions include child teams.

| Type    | Description                                                                                                                                |
| ------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `array` | The teams that can have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos set-user-access-restrictions

https://developer.github.com/v3/repos/branches/#set-user-access-restrictions

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Replaces the list of people that have push access to this branch. This removes all people that previously had push access and grants push access to the new list of people.

| Type    | Description                                                                                                                   |
| ------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos test-push-webhook

https://developer.github.com/v3/repos/hooks/#test-the-push-repository-webhook

This will trigger the hook with the latest push to the current repository if the hook is subscribed to `push` events. If the hook is not subscribed to `push` events, the server will respond with 204 but no test POST will be generated.

**Note**: Previously `/repos/:owner/:repo/hooks/:hook_id/test`

### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |

## repos transfer

https://developer.github.com/v3/repos/#transfer-a-repository

A transfer request will need to be accepted by the new owner when transferring a personal repository to another user. The response will contain the original `owner`, and the transfer will continue asynchronously. For more details on the requirements to transfer personal and organization-owned repositories, see [about repository transfers](https://help.github.com/articles/about-repository-transfers/).

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| new_owner | **Required:** The username or organization name the repository will be transferred to. |
| team_ids | ID of the team or teams to add to the repository. Teams can only be added to organization-owned repositories. |

## repos update

https://developer.github.com/v3/repos/#update-a-repository

**Note**: To edit a repository's topics, use the [Replace all repository topics](https://developer.github.com/v3/repos/#replace-all-repository-topics) endpoint.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| allow_merge_commit | Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits. |
| allow_rebase_merge | Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging. |
| allow_squash_merge | Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging. |
| archived | `true` to archive this repository. **Note**: You cannot unarchive repositories through the API. |
| baptiste-preview | The `is_template` and `template_repository` keys are currently available for developer to preview. See [Create a repository using a template](https://developer.github.com/v3/repos/#create-a-repository-using-a-template) to learn how to create template repositories. To access these new response keys during the preview period, you must set this flag. |
| default_branch | Updates the default branch for this repository. |
| delete_branch_on_merge | Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion. |
| description | A short description of the repository. |
| has_issues | Either `true` to enable issues for this repository or `false` to disable them. |
| has_projects | Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error. |
| has_wiki | Either `true` to enable the wiki for this repository or `false` to disable it. |
| homepage | A URL with more information about the repository. |
| is_template | Either `true` to make this repo available as a template repository or `false` to prevent it. |
| name | The name of the repository. |
| nebula-preview | You can set the visibility of a repository using the new `visibility` parameter in the [Repositories API](https://developer.github.com/v3/repos/), and get a repository's visibility with a new response key. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes/).<br><br>To access repository visibility during the preview period, you must set this flag. |
| private | Either `true` to make the repository private or `false` to make it public. Default: `false`.  <br>**Note**: You will get a `422` error if the organization restricts [changing repository visibility](https://help.github.com/articles/repository-permission-levels-for-an-organization#changing-the-visibility-of-repositories) to organization owners and a non-owner tries to change the value of private. **Note**: You will get a `422` error if the organization restricts [changing repository visibility](https://help.github.com/articles/repository-permission-levels-for-an-organization#changing-the-visibility-of-repositories) to organization owners and a non-owner tries to change the value of private. |
| visibility | Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, `visibility` can also be `internal`. The `visibility` parameter overrides the `private` parameter when you use both along with the `nebula-preview` preview header. |

## repos update-branch-protection

https://developer.github.com/v3/repos/branches/#update-branch-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Protecting a branch requires admin or owner permissions to the repository.

**Note**: Passing new arrays of `users` and `teams` replaces their previous values.

**Note**: The list of users, apps, and teams in total is limited to 100 items.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| enforce_admins | __Required__ Enforce all configured restrictions for administrators. Set to `true` to enforce required status checks for repository administrators. Set to `null` to disable. |
| repo | __Required__ repository in OWNER/REPO form |
| required_status_checks.contexts | __Required__ The list of status checks to require in order to merge into this branch |
| required_status_checks.strict | __Required__ Require branches to be up to date before merging. |
| restrictions.teams | __Required__ The list of team `slug`s with push access |
| restrictions.users | __Required__ The list of user `login`s with push access |
| allow_deletions | Allows deletion of the protected branch by anyone with write access to the repository. Set to `false` to prevent deletion of the protected branch. Default: `false`. For more information, see "[Enabling force pushes to a protected branch](https://help.github.com/en/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)" in the GitHub Help documentation. |
| allow_force_pushes | Permits force pushes to the protected branch by anyone with write access to the repository. Set to `true` to allow force pushes. Set to `false` or `null` to block force pushes. Default: `false`. For more information, see "[Enabling force pushes to a protected branch](https://help.github.com/en/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)" in the GitHub Help documentation." |
| luke-cage-preview | The Protected Branches API now has a setting for requiring a specified number of approving pull request reviews before merging. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-03-16-protected-branches-required-approving-reviews) for full details. To access the API during the preview period, you must set this flag. |
| required_linear_history | Enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch. Set to `true` to enforce a linear commit history. Set to `false` to disable a linear commit Git history. Your repository must allow squash merging or rebase merging before you can enable a linear commit history. Default: `false`. For more information, see "[Requiring a linear commit history](https://help.github.com/github/administering-a-repository/requiring-a-linear-commit-history)" in the GitHub Help documentation. |
| required_pull_request_reviews.dismiss_stale_reviews | Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit. |
| required_pull_request_reviews.dismissal_restrictions.teams | The list of team `slug`s with dismissal access |
| required_pull_request_reviews.dismissal_restrictions.users | The list of user `login`s with dismissal access |
| required_pull_request_reviews.require_code_owner_reviews | Blocks merging pull requests until [code owners](https://help.github.com/articles/about-code-owners/) review them. |
| required_pull_request_reviews.required_approving_review_count | Specify the number of reviewers required to approve pull requests. Use a number between 1 and 6. |
| restrictions.apps | The list of app `slug`s with push access |

## repos update-commit-comment

https://developer.github.com/v3/repos/comments/#update-a-commit-comment



### parameters


| name | description |
|------|-------------|
| body | __Required__ The contents of the comment |
| comment_id | __Required__ comment_id parameter |
| repo | __Required__ repository in OWNER/REPO form |

## repos update-information-about-pages-site

https://developer.github.com/v3/repos/pages/#update-information-about-a-github-pages-site



### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| cname | Specify a custom domain for the repository. Sending a `null` value will remove the custom domain. For more about custom domains, see "[Using a custom domain with GitHub Pages](https://help.github.com/articles/using-a-custom-domain-with-github-pages/)." |
| source |  |

## repos update-invitation

https://developer.github.com/v3/repos/invitations/#update-a-repository-invitation



### parameters


| name | description |
|------|-------------|
| invitation_id | __Required__ invitation_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| permissions | The permissions that the associated user will have on the repository. Valid values are `read`, `write`, `maintain`, `triage`, and `admin`. |

## repos update-pull-request-review-protection

https://developer.github.com/v3/repos/branches/#update-pull-request-review-protection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Updating pull request review enforcement requires admin or owner permissions to the repository and branch protection to be enabled.

**Note**: Passing new arrays of `users` and `teams` replaces their previous values.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| dismiss_stale_reviews | Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit. |
| dismissal_restrictions.teams | The list of team `slug`s with dismissal access |
| dismissal_restrictions.users | The list of user `login`s with dismissal access |
| luke-cage-preview | The Protected Branches API now has a setting for requiring a specified number of approving pull request reviews before merging. This feature is currently available for developers to preview. See the [blog post](https://developer.github.com/changes/2018-03-16-protected-branches-required-approving-reviews) for full details. To access the API during the preview period, you must set this flag. |
| require_code_owner_reviews | Blocks merging pull requests until [code owners](https://help.github.com/articles/about-code-owners/) have reviewed. |
| required_approving_review_count | Specifies the number of reviewers required to approve pull requests. Use a number between 1 and 6. |

## repos update-release

https://developer.github.com/v3/repos/releases/#update-a-release

Users with push access to the repository can edit a release.

### parameters


| name | description |
|------|-------------|
| release_id | __Required__ release_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| body | Text describing the contents of the tag. |
| draft | `true` makes the release a draft, and `false` publishes the release. |
| name | The name of the release. |
| prerelease | `true` to identify the release as a prerelease, `false` to identify the release as a full release. |
| tag_name | The name of the tag. |
| target_commitish | Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch (usually `master`). |

## repos update-release-asset

https://developer.github.com/v3/repos/releases/#update-a-release-asset

Users with push access to the repository can edit a release asset.

### parameters


| name | description |
|------|-------------|
| asset_id | __Required__ asset_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| label | An alternate short description of the asset. Used in place of the filename. |
| name | The file name of the asset. |
| state |  |

## repos update-status-check-protection

https://developer.github.com/v3/repos/branches/#update-status-check-potection

Protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Updating required status checks requires admin or owner permissions to the repository and branch protection to be enabled.

### parameters


| name | description |
|------|-------------|
| branch | __Required__ branch+ parameter |
| repo | __Required__ repository in OWNER/REPO form |
| contexts | The list of status checks to require in order to merge into this branch |
| strict | Require branches to be up to date before merging. |

## repos update-webhook

https://developer.github.com/v3/repos/hooks/#update-a-repository-webhook



### parameters


| name | description |
|------|-------------|
| hook_id | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| active | Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications. |
| add_events | Determines a list of events to be added to the list of events that the Hook triggers for. |
| config.address |  |
| config.content_type | The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`. |
| config.insecure_ssl | Determines whether the SSL certificate of the host for `url` will be verified when delivering payloads. Supported values include `0` (verification is performed) and `1` (verification is not performed). The default is `0`. **We strongly recommend not setting this to `1` as you are subject to man-in-the-middle and other attacks.** |
| config.room |  |
| config.secret | If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value in the [`X-Hub-Signature`](https://developer.github.com/webhooks/event-payloads/#delivery-headers) header. |
| config.url | The URL to which the payloads will be delivered. |
| events | Determines what [events](https://developer.github.com/webhooks/event-payloads) the hook is triggered for. This replaces the entire array of events. |
| remove_events | Determines a list of events to be removed from the list of events that the Hook triggers for. |

## repos upload-release-asset

https://developer.github.com/v3/repos/releases/#upload-a-release-asset

This endpoint makes use of [a Hypermedia relation](https://developer.github.com/v3/#hypermedia) to determine which URL to access. The endpoint you call to upload release assets is specific to your release. Use the `upload_url` returned in
the response of the [Create a release endpoint](https://developer.github.com/v3/repos/releases/#create-a-release) to upload a release asset.

You need to use an HTTP client which supports [SNI](http://en.wikipedia.org/wiki/Server_Name_Indication) to make calls to this endpoint.

Most libraries will set the required `Content-Length` header automatically. Use the required `Content-Type` header to provide the media type of the asset. For a list of media types, see [Media Types](https://www.iana.org/assignments/media-types/media-types.xhtml). For example: 

`application/zip`

GitHub expects the asset data in its raw binary form, rather than JSON. You will send the raw binary content of the asset as the request body. Everything else about the endpoint is the same as the rest of the API. For example,
you'll still need to pass your authentication to be able to upload an asset.

When an upstream failure occurs, you will receive a `502 Bad Gateway` status. This may leave an empty asset with a state of `starter`. It can be safely deleted.

**Notes:**
*   GitHub renames asset filenames that have special characters, non-alphanumeric characters, and leading or trailing periods. The "[List assets for a release](https://developer.github.com/v3/repos/releases/#list-assets-for-a-release)"
endpoint lists the renamed filenames. For more information and help, contact [GitHub Support](https://github.com/contact).
*   If you upload an asset with the same filename as another uploaded asset, you'll receive an error and must delete the old file before you can re-upload the new asset.

### parameters


| name | description |
|------|-------------|
| file | __Required__ the file to upload |
| release_id | __Required__ release_id parameter |
| repo | __Required__ repository in OWNER/REPO form |
| content-type | override the Content-Type header |
| label | label parameter |
| name | name parameter |

# scim


## scim delete-user-from-org

https://developer.github.com/v3/scim/#delete-a-scim-user-from-an-organization



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| scim_user_id | __Required__ scim_user_id parameter |

## scim get-provisioning-information-for-user

https://developer.github.com/v3/scim/#get-scim-provisioning-information-for-a-user



### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| scim_user_id | __Required__ scim_user_id parameter |

## scim list-provisioned-identities

https://developer.github.com/v3/scim/#list-scim-provisioned-identities

Retrieves a paginated list of all provisioned organization members, including pending invitations. If you provide the `filter` parameter, the resources for all matching provisions members are returned.

When a user with a SAML-provisioned external identity leaves (or is removed from) an organization, the account's metadata is immediately removed. However, the returned list of user accounts might not always match the organization or enterprise member list you see on GitHub. This can happen in certain cases where an external identity associated with an organization will not match an organization member:
  - When a user with a SCIM-provisioned external identity is removed from an organization, the account's metadata is preserved to allow the user to re-join the organization in the future.
  - When inviting a user to join an organization, you can expect to see their external identity in the results before they accept the invitation, or if the invitation is cancelled (or never accepted).
  - When a user is invited over SCIM, an external identity is created that matches with the invitee's email address. However, this identity is only linked to a user account when the user accepts the invitation by going through SAML SSO.

The returned list of external identities can include an entry for a `null` user. These are unlinked SAML identities that are created when a user goes through the following Single Sign-On (SSO) process but does not sign in to their GitHub account after completing SSO:

1. The user is granted access by the IdP and is not a member of the GitHub organization.

1. The user attempts to access the GitHub organization and initiates the SAML SSO process, and is not currently signed in to their GitHub account.

1. After successfully authenticating with the SAML SSO IdP, the `null` external identity entry is created and the user is prompted to sign in to their GitHub account:
   - If the user signs in, their GitHub account is linked to this entry.
   - If the user does not sign in (or does not create a new account when prompted), they are not added to the GitHub organization, and the external identity `null` entry remains in place.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| count | Used for pagination: the number of results to return. |
| filter | Filters results using the equals query parameter operator (`eq`). You can filter results that are equal to `id`, `userName`, `emails`, and `external_id`. For example, to search for an identity with the `userName` Octocat, you would use this query:<br><br>`?filter=userName%20eq%20\"Octocat\"`.<br><br>To filter results for for the identity with the email `octocat@github.com`, you would use this query:<br><br>`?filter=emails%20eq%20\"octocat@github.com\"`. |
| startIndex | Used for pagination: the index of the first result to return. |

## scim provision-and-invite-user

https://developer.github.com/v3/scim/#provision-and-invite-a-scim-user

Provision organization membership for a user, and send an activation email to the email address.

### parameters


| name | description |
|------|-------------|
| emails | __Required__ user emails |
| name.familyName | __Required__  |
| name.givenName | __Required__  |
| org | __Required__  |
| userName | __Required__ Configured by the admin. Could be an email, login, or username |
| active |  |
| displayName | The name of the user, suitable for display to end-users |
| externalId |  |
| groups |  |
| name.formatted |  |
| schemas |  |

## scim set-information-for-provisioned-user

https://developer.github.com/v3/scim/#set-scim-information-for-a-provisioned-user

Replaces an existing provisioned user's information. You must provide all the information required for the user as if you were provisioning them for the first time. Any existing user information that you don't provide will be removed. If you want to only update a specific attribute, use the [Update an attribute for a SCIM user](https://developer.github.com/v3/scim/#update-an-attribute-for-a-scim-user) endpoint instead.

You must at least provide the required values for the user: `userName`, `name`, and `emails`.

**Warning:** Setting `active: false` removes the user from the organization, deletes the external identity, and deletes the associated `{scim_user_id}`.

### parameters


| name | description |
|------|-------------|
| emails | __Required__ user emails |
| name.familyName | __Required__  |
| name.givenName | __Required__  |
| org | __Required__  |
| scim_user_id | __Required__ scim_user_id parameter |
| userName | __Required__ Configured by the admin. Could be an email, login, or username |
| active |  |
| displayName | The name of the user, suitable for display to end-users |
| externalId |  |
| groups |  |
| name.formatted |  |
| schemas |  |

## scim update-attribute-for-user

https://developer.github.com/v3/scim/#update-an-attribute-for-a-scim-user

Allows you to change a provisioned user's individual attributes. To change a user's values, you must provide a specific `Operations` JSON format that contains at least one of the `add`, `remove`, or `replace` operations. For examples and more information on the SCIM operations format, see the [SCIM specification](https://tools.ietf.org/html/rfc7644#section-3.5.2).

**Note:** Complicated SCIM `path` selectors that include filters are not supported. For example, a `path` selector defined as `"path": "emails[type eq \"work\"]"` will not work.

**Warning:** If you set `active:false` using the `replace` operation (as shown in the JSON example below), it removes the user from the organization, deletes the external identity, and deletes the associated `:scim_user_id`.

```
{
  "Operations":[{
    "op":"replace",
    "value":{
      "active":false
    }
  }]
}
```

### parameters


| name | description |
|------|-------------|
| Operations | __Required__ Set of operations to be performed |
| org | __Required__  |
| scim_user_id | __Required__ scim_user_id parameter |
| schemas |  |

# search


## search code

https://developer.github.com/v3/search/#search-code

Searches for query terms inside of a file. This method returns up to 100 results [per page](https://developer.github.com/v3/#pagination).

When searching for code, you can get text match metadata for the file **content** and file **path** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://developer.github.com/v3/search/#text-match-metadata).

For example, if you want to find the definition of the `addClass` function inside [jQuery](https://github.com/jquery/jquery) repository, your query would look something like this:

`q=addClass+in:file+language:js+repo:jquery/jquery`

This query searches for the keyword `addClass` within a file's contents. The query limits the search to files where the language is JavaScript in the `jquery/jquery` repository.

#### Considerations for code search

Due to the complexity of searching code, there are a few restrictions on how searches are performed:

*   Only the _default branch_ is considered. In most cases, this will be the `master` branch.
*   Only files smaller than 384 KB are searchable.
*   You must always include at least one search term when searching source code. For example, searching for [`language:go`](https://github.com/search?utf8=%E2%9C%93&q=language%3Ago&type=Code) is not valid, while [`amazing
language:go`](https://github.com/search?utf8=%E2%9C%93&q=amazing+language%3Ago&type=Code) is.

### parameters


| name | description |
|------|-------------|
| q | __Required__ The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as GitHub.com. To learn more about the format of the query, see [Constructing a search query](https://developer.github.com/v3/search/#constructing-a-search-query). See "[Searching code](https://help.github.com/articles/searching-code/)" for a detailed list of qualifiers. |
| order | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | Sorts the results of your query. Can only be `indexed`, which indicates how recently a file has been indexed by the GitHub search infrastructure. Default: [best match](https://developer.github.com/v3/search/#ranking-search-results) |

## search commits

https://developer.github.com/v3/search/#search-commits

Find commits via various criteria on the default branch (usually `master`). This method returns up to 100 results [per page](https://developer.github.com/v3/#pagination).

When searching for commits, you can get text match metadata for the **message** field when you provide the `text-match` media type. For more details about how to receive highlighted search results, see [Text match
metadata](https://developer.github.com/v3/search/#text-match-metadata).

For example, if you want to find commits related to CSS in the [octocat/Spoon-Knife](https://github.com/octocat/Spoon-Knife) repository. Your query would look something like this:

`q=repo:octocat/Spoon-Knife+css`

### parameters


| name | description |
|------|-------------|
| cloak-preview | __Required__ The Commit Search API is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2017-01-05-commit-search-api/) for full details.<br><br>To access the API you must set this flag. |
| q | __Required__ The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as GitHub.com. To learn more about the format of the query, see [Constructing a search query](https://developer.github.com/v3/search/#constructing-a-search-query). See "[Searching commits](https://help.github.com/articles/searching-commits/)" for a detailed list of qualifiers. |
| order | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | Sorts the results of your query by `author-date` or `committer-date`. Default: [best match](https://developer.github.com/v3/search/#ranking-search-results) |

## search issues-and-pull-requests

https://developer.github.com/v3/search/#search-issues-and-pull-requests

Find issues by state and keyword. This method returns up to 100 results [per page](https://developer.github.com/v3/#pagination).

When searching for issues, you can get text match metadata for the issue **title**, issue **body**, and issue **comment body** fields when you pass the `text-match` media type. For more details about how to receive highlighted
search results, see [Text match metadata](https://developer.github.com/v3/search/#text-match-metadata).

For example, if you want to find the oldest unresolved Python bugs on Windows. Your query might look something like this.

`q=windows+label:bug+language:python+state:open&sort=created&order=asc`

This query searches for the keyword `windows`, within any open issue that is labeled as `bug`. The search runs across repositories whose primary language is Python. The results are sorted by creation date in ascending order, whick means the oldest issues appear first in the search results.

### parameters


| name | description |
|------|-------------|
| q | __Required__ The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as GitHub.com. To learn more about the format of the query, see [Constructing a search query](https://developer.github.com/v3/search/#constructing-a-search-query). See "[Searching issues and pull requests](https://help.github.com/articles/searching-issues-and-pull-requests/)" for a detailed list of qualifiers. |
| order | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | Sorts the results of your query by the number of `comments`, `reactions`, `reactions-+1`, `reactions--1`, `reactions-smile`, `reactions-thinking_face`, `reactions-heart`, `reactions-tada`, or `interactions`. You can also sort results by how recently the items were `created` or `updated`, Default: [best match](https://developer.github.com/v3/search/#ranking-search-results) |

## search labels

https://developer.github.com/v3/search/#search-labels

Find labels in a repository with names or descriptions that match search keywords. Returns up to 100 results [per page](https://developer.github.com/v3/#pagination).

When searching for labels, you can get text match metadata for the label **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://developer.github.com/v3/search/#text-match-metadata).

For example, if you want to find labels in the `linguist` repository that match `bug`, `defect`, or `enhancement`. Your query might look like this:

`q=bug+defect+enhancement&repository_id=64778136`

The labels that best match the query appear first in the search results.

### parameters


| name | description |
|------|-------------|
| q | __Required__ The search keywords. This endpoint does not accept qualifiers in the query. To learn more about the format of the query, see [Constructing a search query](https://developer.github.com/v3/search/#constructing-a-search-query). |
| repository_id | __Required__ The id of the repository. |
| order | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. |
| sort | Sorts the results of your query by when the label was `created` or `updated`. Default: [best match](https://developer.github.com/v3/search/#ranking-search-results) |

## search repos

https://developer.github.com/v3/search/#search-repositories

Find repositories via various criteria. This method returns up to 100 results [per page](https://developer.github.com/v3/#pagination).

When searching for repositories, you can get text match metadata for the **name** and **description** fields when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://developer.github.com/v3/search/#text-match-metadata).

For example, if you want to search for popular Tetris repositories written in assembly code, your query might look like this:

`q=tetris+language:assembly&sort=stars&order=desc`

This query searches for repositories with the word `tetris` in the name, the description, or the README. The results are limited to repositories where the primary language is assembly. The results are sorted by stars in descending order, so that the most popular repositories appear first in the search results.

When you include the `mercy` preview header, you can also search for multiple topics by adding more `topic:` instances. For example, your query might look like this:

`q=topic:ruby+topic:rails`

### parameters


| name | description |
|------|-------------|
| q | __Required__ The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as GitHub.com. To learn more about the format of the query, see [Constructing a search query](https://developer.github.com/v3/search/#constructing-a-search-query). See "[Searching for repositories](https://help.github.com/articles/searching-for-repositories/)" for a detailed list of qualifiers. |
| mercy-preview | The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must set this flag. |
| order | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | Sorts the results of your query by number of `stars`, `forks`, or `help-wanted-issues` or how recently the items were `updated`. Default: [best match](https://developer.github.com/v3/search/#ranking-search-results) |

## search topics

https://developer.github.com/v3/search/#search-topics

Find topics via various criteria. Results are sorted by best match. This method returns up to 100 results [per page](https://developer.github.com/v3/#pagination). See "[Searching topics](https://help.github.com/articles/searching-topics/)" for a detailed list of qualifiers.

When searching for topics, you can get text match metadata for the topic's **short\_description**, **description**, **name**, or **display\_name** field when you pass the `text-match` media type. For more details about how to receive highlighted search results, see [Text match metadata](https://developer.github.com/v3/search/#text-match-metadata).

For example, if you want to search for topics related to Ruby that are featured on https://github.com/topics. Your query might look like this:

`q=ruby+is:featured`

This query searches for topics with the keyword `ruby` and limits the results to find only topics that are featured. The topics that are the best match for the query appear first in the search results.

### parameters


| name | description |
|------|-------------|
| mercy-preview | __Required__ The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must set this flag. |
| q | __Required__ The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as GitHub.com. To learn more about the format of the query, see [Constructing a search query](https://developer.github.com/v3/search/#constructing-a-search-query). |

## search users

https://developer.github.com/v3/search/#search-users

Find users via various criteria. This method returns up to 100 results [per page](https://developer.github.com/v3/#pagination).

When searching for users, you can get text match metadata for the issue **login**, **email**, and **name** fields when you pass the `text-match` media type. For more details about highlighting search results, see [Text match metadata](https://developer.github.com/v3/search/#text-match-metadata). For more details about how to receive highlighted search results, see [Text match metadata](https://developer.github.com/v3/search/#text-match-metadata).

For example, if you're looking for a list of popular users, you might try this query:

`q=tom+repos:%3E42+followers:%3E1000`

This query searches for users with the name `tom`. The results are restricted to users with more than 42 repositories and over 1,000 followers.

### parameters


| name | description |
|------|-------------|
| q | __Required__ The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as GitHub.com. To learn more about the format of the query, see [Constructing a search query](https://developer.github.com/v3/search/#constructing-a-search-query). See "[Searching users](https://help.github.com/articles/searching-users/)" for a detailed list of qualifiers. |
| order | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| sort | Sorts the results of your query by number of `followers` or `repositories`, or when the person `joined` GitHub. Default: [best match](https://developer.github.com/v3/search/#ranking-search-results) |

# teams


## teams add-member-legacy

https://developer.github.com/v3/teams/members/#add-team-member-legacy

The "Add team member" endpoint (described below) is deprecated.

We recommend using the [Add or update team membership for a user](https://developer.github.com/v3/teams/members/#add-or-update-team-membership-for-a-user) endpoint instead. It allows you to invite new organization members to your teams.

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

To add someone to a team, the authenticated user must be an organization owner or a team maintainer in the team they're changing. The person being added to the team must be a member of the team's organization.

**Note:** When you have team synchronization set up for a team with your organization's identity provider (IdP), you will see an error if you attempt to use the API for making changes to the team's membership. If you have access to manage group membership in your IdP, you can manage GitHub team membership through your identity provider, which automatically adds and removes team members in an organization. For more information, see "[Synchronizing teams between your identity provider and GitHub](https://help.github.com/articles/synchronizing-teams-between-your-identity-provider-and-github/)."

Note that you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| username | __Required__  |

## teams add-or-update-membership-for-user-in-org

https://developer.github.com/v3/teams/members/#add-or-update-team-membership-for-a-user

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Adds an organization member to a team. An authenticated organization owner or team maintainer can add organization members to a team.

**Note:** When you have team synchronization set up for a team with your organization's identity provider (IdP), you will see an error if you attempt to use the API for making changes to the team's membership. If you have access to manage group membership in your IdP, you can manage GitHub team membership through your identity provider, which automatically adds and removes team members in an organization. For more information, see "[Synchronizing teams between your identity provider and GitHub](https://help.github.com/articles/synchronizing-teams-between-your-identity-provider-and-github/)."

An organization owner can add someone who is not part of the team's organization to a team. When an organization owner adds someone to a team who is not an organization member, this endpoint will send an invitation to the person via email. This newly-created membership will be in the "pending" state until the person accepts the invitation, at which point the membership will transition to the "active" state and the user will be added as a member of the team.

If the user is already a member of the team, this endpoint will update the role of the team member's role. To update the membership of a team member, the authenticated user must be an organization owner or a team maintainer.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `PUT /organizations/{org_id}/team/{team_id}/memberships/{username}`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| username | __Required__  |
| role | The role that this user should have in the team. Can be one of:  <br>\* `member` - a normal member of the team.  <br>\* `maintainer` - a team maintainer. Able to add/remove other team members, promote other team members to team maintainer, and edit the team's name and description. |

## teams add-or-update-membership-for-user-legacy

https://developer.github.com/v3/teams/members/#add-or-update-team-membership-for-a-user-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Add or update team membership for a user](https://developer.github.com/v3/teams/members/#add-or-update-team-membership-for-a-user) endpoint.

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

If the user is already a member of the team's organization, this endpoint will add the user to the team. To add a membership between an organization member and a team, the authenticated user must be an organization owner or a team maintainer.

**Note:** When you have team synchronization set up for a team with your organization's identity provider (IdP), you will see an error if you attempt to use the API for making changes to the team's membership. If you have access to manage group membership in your IdP, you can manage GitHub team membership through your identity provider, which automatically adds and removes team members in an organization. For more information, see "[Synchronizing teams between your identity provider and GitHub](https://help.github.com/articles/synchronizing-teams-between-your-identity-provider-and-github/)."

If the user is unaffiliated with the team's organization, this endpoint will send an invitation to the user via email. This newly-created membership will be in the "pending" state until the user accepts the invitation, at which point the membership will transition to the "active" state and the user will be added as a member of the team. To add a membership between an unaffiliated user and a team, the authenticated user must be an organization owner.

If the user is already a member of the team, this endpoint will update the role of the team member's role. To update the membership of a team member, the authenticated user must be an organization owner or a team maintainer.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| username | __Required__  |
| role | The role that this user should have in the team. Can be one of:  <br>\* `member` - a normal member of the team.  <br>\* `maintainer` - a team maintainer. Able to add/remove other team members, promote other team members to team maintainer, and edit the team's name and description. |

## teams add-or-update-project-permissions-in-org

https://developer.github.com/v3/teams/#add-or-update-team-project-permissions

Adds an organization project to a team. To add a project to a team or update the team's permission on a project, the authenticated user must have `admin` permissions for the project. The project and team must be part of the same organization.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `PUT /organizations/{org_id}/team/{team_id}/projects/{project_id}`.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| org | __Required__  |
| project_id | __Required__  |
| team_slug | __Required__ team_slug parameter |
| permission | The permission to grant to the team for this project. Can be one of:  <br>\* `read` - team members can read, but not write to or administer this project.  <br>\* `write` - team members can read and write, but not administer this project.  <br>\* `admin` - team members can read, write and administer this project.  <br>Default: the team's `permission` attribute will be used to determine what permission to grant the team on this project. Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)." |

## teams add-or-update-project-permissions-legacy

https://developer.github.com/v3/teams/#add-or-update-team-project-permissions-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Add or update team project permissions](https://developer.github.com/v3/teams/#add-or-update-team-project-permissions) endpoint.

Adds an organization project to a team. To add a project to a team or update the team's permission on a project, the authenticated user must have `admin` permissions for the project. The project and team must be part of the same organization.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |
| team_id | __Required__  |
| permission | The permission to grant to the team for this project. Can be one of:  <br>\* `read` - team members can read, but not write to or administer this project.  <br>\* `write` - team members can read and write, but not administer this project.  <br>\* `admin` - team members can read, write and administer this project.  <br>Default: the team's `permission` attribute will be used to determine what permission to grant the team on this project. Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)." |

## teams add-or-update-repo-permissions-in-org

https://developer.github.com/v3/teams/#add-or-update-team-repository-permissions

To add a repository to a team or update the team's permission on a repository, the authenticated user must have admin access to the repository, and must be able to see the team. The repository must be owned by the organization, or a direct fork of a repository owned by the organization. You will get a `422 Unprocessable Entity` status if you attempt to add a repository to a team that is not owned by the organization. Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

**Note:** You can also specify a team by `org_id` and `team_id` using the route `PUT /organizations/{org_id}/team/{team_id}/repos/{owner}/{repo}`.

For more information about the permission levels, see "[Repository permission levels for an organization](https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/repository-permission-levels-for-an-organization#permission-levels-for-repositories-owned-by-an-organization)".

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| team_slug | __Required__ team_slug parameter |
| permission | The permission to grant the team on this repository. Can be one of:  <br>\* `pull` - team members can pull, but not push to or administer this repository.  <br>\* `push` - team members can pull and push, but not administer this repository.  <br>\* `admin` - team members can pull, push and administer this repository.  <br>\* `maintain` - team members can manage the repository without access to sensitive or destructive actions. Recommended for project managers. Only applies to repositories owned by organizations.  <br>\* `triage` - team members can proactively manage issues and pull requests without write access. Recommended for contributors who triage a repository. Only applies to repositories owned by organizations.  <br>  <br>If no permission is specified, the team's `permission` attribute will be used to determine what permission to grant the team on this repository. |

## teams add-or-update-repo-permissions-legacy

https://developer.github.com/v3/teams/#add-or-update-team-repository-permissions-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Add or update team repository permissions](https://developer.github.com/v3/teams/#add-or-update-team-repository-permissions) endpoint.

To add a repository to a team or update the team's permission on a repository, the authenticated user must have admin access to the repository, and must be able to see the team. The repository must be owned by the organization, or a direct fork of a repository owned by the organization. You will get a `422 Unprocessable Entity` status if you attempt to add a repository to a team that is not owned by the organization.

Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| team_id | __Required__  |
| permission | The permission to grant the team on this repository. Can be one of:  <br>\* `pull` - team members can pull, but not push to or administer this repository.  <br>\* `push` - team members can pull and push, but not administer this repository.  <br>\* `admin` - team members can pull, push and administer this repository.  <br>  <br>If no permission is specified, the team's `permission` attribute will be used to determine what permission to grant the team on this repository. |

## teams check-permissions-for-project-in-org

https://developer.github.com/v3/teams/#check-team-permissions-for-a-project

Checks whether a team has `read`, `write`, or `admin` permissions for an organization project. The response includes projects inherited from a parent team.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/projects/{project_id}`.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| org | __Required__  |
| project_id | __Required__  |
| team_slug | __Required__ team_slug parameter |

## teams check-permissions-for-project-legacy

https://developer.github.com/v3/teams/#check-team-permissions-for-a-project-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Check team permissions for a project](https://developer.github.com/v3/teams/#check-team-permissions-for-a-project) endpoint.

Checks whether a team has `read`, `write`, or `admin` permissions for an organization project. The response includes projects inherited from a parent team.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| project_id | __Required__  |
| team_id | __Required__  |

## teams check-permissions-for-repo-in-org

https://developer.github.com/v3/teams/#check-team-permissions-for-a-repository

Checks whether a team has `admin`, `push`, `maintain`, `triage`, or `pull` permission for a repository. Repositories inherited through a parent team will also be checked.

You can also get information about the specified repository, including what permissions the team grants on it, by passing the following custom [media type](https://developer.github.com/v3/media/) via the `application/vnd.github.v3.repository+json` accept header.

If a team doesn't have permission for the repository, you will receive a `404 Not Found` response status.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/repos/{owner}/{repo}`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| team_slug | __Required__ team_slug parameter |

## teams check-permissions-for-repo-legacy

https://developer.github.com/v3/teams/#check-team-permissions-for-a-repository-legacy

**Note**: Repositories inherited through a parent team will also be checked.

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Check team permissions for a repository](https://developer.github.com/v3/teams/#check-team-permissions-for-a-repository) endpoint.

You can also get information about the specified repository, including what permissions the team grants on it, by passing the following custom [media type](https://developer.github.com/v3/media/) via the `Accept` header:

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| team_id | __Required__  |

## teams create

https://developer.github.com/v3/teams/#create-a-team

To create a team, the authenticated user must be a member or owner of `{org}`. By default, organization members can create teams. Organization owners can limit team creation to organization owners. For more information, see "[Setting team creation permissions](https://help.github.com/en/articles/setting-team-creation-permissions-in-your-organization)."

When you create a new team, you automatically become a team maintainer without explicitly adding yourself to the optional array of `maintainers`. For more information, see "[About teams](https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/about-teams)".

### parameters


| name | description |
|------|-------------|
| name | __Required__ The name of the team. |
| org | __Required__  |
| description | The description of the team. |
| maintainers | List GitHub IDs for organization members who will become team maintainers. |
| parent_team_id | The ID of a team to set as the parent team. |
| permission | **Deprecated**. The permission that new repositories will be added to the team with when none is specified. Can be one of:  <br>\* `pull` - team members can pull, but not push to or administer newly-added repositories.  <br>\* `push` - team members can pull and push, but not administer newly-added repositories.  <br>\* `admin` - team members can pull, push and administer newly-added repositories. |
| privacy | The level of privacy this team should have. The options are:  <br>**For a non-nested team:**  <br>\* `secret` - only visible to organization owners and members of this team.  <br>\* `closed` - visible to all members of this organization.  <br>Default: `secret`  <br>**For a parent or child team:**  <br>\* `closed` - visible to all members of this organization.  <br>Default for child team: `closed` |
| repo_names | The full name (e.g., "organization-name/repository-name") of repositories to add the team to. |

## teams create-discussion-comment-in-org

https://developer.github.com/v3/teams/discussion_comments/#create-a-discussion-comment

Creates a new comment on a team discussion. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `POST /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments`.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The discussion comment's body text. |
| discussion_number | __Required__  |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams create-discussion-comment-legacy

https://developer.github.com/v3/teams/discussion_comments/#create-a-discussion-comment-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Create a discussion comment](https://developer.github.com/v3/teams/discussion_comments/#create-a-discussion-comment) endpoint.

Creates a new comment on a team discussion. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The discussion comment's body text. |
| discussion_number | __Required__  |
| team_id | __Required__  |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams create-discussion-in-org

https://developer.github.com/v3/teams/discussions/#create-a-discussion

Creates a new discussion post on a team's page. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `POST /organizations/{org_id}/team/{team_id}/discussions`.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The discussion post's body text. |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| title | __Required__ The discussion post's title. |
| private | Private posts are only visible to team members, organization owners, and team maintainers. Public posts are visible to all members of the organization. Set to `true` to create a private post. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams create-discussion-legacy

https://developer.github.com/v3/teams/discussions/#create-a-discussion-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Create a discussion`](https://developer.github.com/v3/teams/discussions/#create-a-discussion) endpoint.

Creates a new discussion post on a team's page. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

This endpoint triggers [notifications](https://help.github.com/articles/about-notifications/). Creating content too quickly using this endpoint may result in abuse rate limiting. See "[Abuse rate limits](https://developer.github.com/v3/#abuse-rate-limits)" and "[Dealing with abuse rate limits](https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits)" for details.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The discussion post's body text. |
| team_id | __Required__  |
| title | __Required__ The discussion post's title. |
| private | Private posts are only visible to team members, organization owners, and team maintainers. Public posts are visible to all members of the organization. Set to `true` to create a private post. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams create-or-update-idp-group-connections-in-org

https://developer.github.com/v3/teams/team_sync/#create-or-update-idp-group-connections

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Creates, updates, or removes a connection between a team and an IdP group. When adding groups to a team, you must include all new and existing groups to avoid replacing existing groups with the new ones. Specifying an empty `groups` array will remove all connections for a team.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `PATCH /organizations/{org_id}/team/{team_id}/team-sync/group-mappings`.

### parameters


| name | description |
|------|-------------|
| groups | __Required__ The IdP groups you want to connect to a GitHub team. When updating, the new `groups` object will replace the original one. You must include any existing groups that you don't want to remove. |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |

## teams create-or-update-idp-group-connections-legacy

https://developer.github.com/v3/teams/team_sync/#create-or-update-idp-group-connections-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Create or update IdP group connections`](https://developer.github.com/v3/teams/team_sync/#create-or-update-idp-group-connections) endpoint.

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

Creates, updates, or removes a connection between a team and an IdP group. When adding groups to a team, you must include all new and existing groups to avoid replacing existing groups with the new ones. Specifying an empty `groups` array will remove all connections for a team.

### parameters


| name | description |
|------|-------------|
| groups | __Required__ The IdP groups you want to connect to a GitHub team. When updating, the new `groups` object will replace the original one. You must include any existing groups that you don't want to remove. |
| team_id | __Required__  |
| synced_at |  |

## teams delete-discussion-comment-in-org

https://developer.github.com/v3/teams/discussion_comments/#delete-a-discussion-comment

Deletes a comment on a team discussion. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| discussion_number | __Required__  |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |

## teams delete-discussion-comment-legacy

https://developer.github.com/v3/teams/discussion_comments/#delete-a-discussion-comment-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Delete a discussion comment](https://developer.github.com/v3/teams/discussion_comments/#delete-a-discussion-comment) endpoint.

Deletes a comment on a team discussion. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| discussion_number | __Required__  |
| team_id | __Required__  |

## teams delete-discussion-in-org

https://developer.github.com/v3/teams/discussions/#delete-a-discussion

Delete a discussion from a team's page. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}`.

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |

## teams delete-discussion-legacy

https://developer.github.com/v3/teams/discussions/#delete-a-discussion-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Delete a discussion`](https://developer.github.com/v3/teams/discussions/#delete-a-discussion) endpoint.

Delete a discussion from a team's page. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| team_id | __Required__  |

## teams delete-in-org

https://developer.github.com/v3/teams/#delete-a-team

To delete a team, the authenticated user must be an organization owner or team maintainer.

If you are an organization owner, deleting a parent team will delete all of its child teams as well.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |

## teams delete-legacy

https://developer.github.com/v3/teams/#delete-a-team-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Delete a team](https://developer.github.com/v3/teams/#delete-a-team) endpoint.

To delete a team, the authenticated user must be an organization owner or team maintainer.

If you are an organization owner, deleting a parent team will delete all of its child teams as well.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |

## teams get-by-name

https://developer.github.com/v3/teams/#get-a-team-by-name

Gets a team using the team's `slug`. GitHub generates the `slug` from the team `name`.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |

## teams get-discussion-comment-in-org

https://developer.github.com/v3/teams/discussion_comments/#get-a-discussion-comment

Get a specific comment on a team discussion. OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| discussion_number | __Required__  |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams get-discussion-comment-legacy

https://developer.github.com/v3/teams/discussion_comments/#get-a-discussion-comment-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Get a discussion comment](https://developer.github.com/v3/teams/discussion_comments/#get-a-discussion-comment) endpoint.

Get a specific comment on a team discussion. OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| comment_number | __Required__  |
| discussion_number | __Required__  |
| team_id | __Required__  |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams get-discussion-in-org

https://developer.github.com/v3/teams/discussions/#get-a-discussion

Get a specific discussion on a team's page. OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}`.

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams get-discussion-legacy

https://developer.github.com/v3/teams/discussions/#get-a-discussion-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Get a discussion](https://developer.github.com/v3/teams/discussions/#get-a-discussion) endpoint.

Get a specific discussion on a team's page. OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| team_id | __Required__  |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams get-legacy

https://developer.github.com/v3/teams/#get-a-team-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the [Get a team by name](https://developer.github.com/v3/teams/#get-a-team-by-name) endpoint.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |

## teams get-member-legacy

https://developer.github.com/v3/teams/members/#get-team-member-legacy

The "Get team member" endpoint (described below) is deprecated.

We recommend using the [Get team membership for a user](https://developer.github.com/v3/teams/members/#get-team-membership-for-a-user) endpoint instead. It allows you to get both active and pending memberships.

To list members in a team, the team must be visible to the authenticated user.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| username | __Required__  |

## teams get-membership-for-user-in-org

https://developer.github.com/v3/teams/members/#get-team-membership-for-a-user

Team members will include the members of child teams.

To get a user's membership with a team, the team must be visible to the authenticated user.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/memberships/{username}`.

**Note:** The `role` for organization owners returns as `maintainer`. For more information about `maintainer` roles, see [Create a team](https://developer.github.com/v3/teams/#create-a-team).

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| username | __Required__  |

## teams get-membership-for-user-legacy

https://developer.github.com/v3/teams/members/#get-team-membership-for-a-user-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Get team membership for a user](https://developer.github.com/v3/teams/members/#get-team-membership-for-a-user) endpoint.

Team members will include the members of child teams.

To get a user's membership with a team, the team must be visible to the authenticated user.

**Note:** The `role` for organization owners returns as `maintainer`. For more information about `maintainer` roles, see [Create a team](https://developer.github.com/v3/teams/#create-a-team).

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| username | __Required__  |

## teams list

https://developer.github.com/v3/teams/#list-teams

Lists all teams in an organization that are visible to the authenticated user.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-child-in-org

https://developer.github.com/v3/teams/#list-child-teams

Lists the child teams of the team specified by `{team_slug}`.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/teams`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-child-legacy

https://developer.github.com/v3/teams/#list-child-teams-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List child teams`](https://developer.github.com/v3/teams/#list-child-teams) endpoint.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-discussion-comments-in-org

https://developer.github.com/v3/teams/discussion_comments/#list-discussion-comments

List all comments on a team discussion. OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments`.

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| direction | One of `asc` (ascending) or `desc` (descending). |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams list-discussion-comments-legacy

https://developer.github.com/v3/teams/discussion_comments/#list-discussion-comments-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [List discussion comments](https://developer.github.com/v3/teams/discussion_comments/#list-discussion-comments) endpoint.

List all comments on a team discussion. OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| team_id | __Required__  |
| direction | One of `asc` (ascending) or `desc` (descending). |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams list-discussions-in-org

https://developer.github.com/v3/teams/discussions/#list-discussions

List all discussions on a team's page. OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/discussions`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| direction | One of `asc` (ascending) or `desc` (descending). |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams list-discussions-legacy

https://developer.github.com/v3/teams/discussions/#list-discussions-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List discussions`](https://developer.github.com/v3/teams/discussions/#list-discussions) endpoint.

List all discussions on a team's page. OAuth access tokens require the `read:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| direction | One of `asc` (ascending) or `desc` (descending). |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams list-for-authenticated-user

https://developer.github.com/v3/teams/#list-teams-for-the-authenticated-user

List all of the teams across all of the organizations to which the authenticated user belongs. This method requires `user`, `repo`, or `read:org` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/) when authenticating via [OAuth](https://developer.github.com/apps/building-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-idp-groups-for-legacy

https://developer.github.com/v3/teams/team_sync/#list-idp-groups-for-a-team-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List IdP groups for a team`](https://developer.github.com/v3/teams/team_sync/#list-idp-groups-for-a-team) endpoint.

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

List IdP groups connected to a team on GitHub.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |

## teams list-idp-groups-for-org

https://developer.github.com/v3/teams/team_sync/#list-idp-groups-for-an-organization

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

List IdP groups available in an organization. You can limit your page results using the `per_page` parameter. GitHub generates a url-encoded `page` token using a cursor value for where the next page begins. For more information on cursor pagination, see "[Offset and Cursor Pagination explained](https://dev.to/jackmarchant/offset-and-cursor-pagination-explained-b89)."

The `per_page` parameter provides pagination for a list of IdP groups the authenticated user can access in an organization. For example, if the user `octocat` wants to see two groups per page in `octo-org` via cURL, it would look like this:

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-idp-groups-in-org

https://developer.github.com/v3/teams/team_sync/#list-idp-groups-for-a-team

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

List IdP groups connected to a team on GitHub.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/team-sync/group-mappings`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |

## teams list-members-in-org

https://developer.github.com/v3/teams/members/#list-team-members

Team members will include the members of child teams.

To list members in a team, the team must be visible to the authenticated user.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| role | Filters members returned by their role in the team. Can be one of:  <br>\* `member` - normal members of the team.  <br>\* `maintainer` - team maintainers.  <br>\* `all` - all members of the team. |

## teams list-members-legacy

https://developer.github.com/v3/teams/members/#list-team-members-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List team members`](https://developer.github.com/v3/teams/members/#list-team-members) endpoint.

Team members will include the members of child teams.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |
| role | Filters members returned by their role in the team. Can be one of:  <br>\* `member` - normal members of the team.  <br>\* `maintainer` - team maintainers.  <br>\* `all` - all members of the team. |

## teams list-pending-invitations-in-org

https://developer.github.com/v3/teams/members/#list-pending-team-invitations

The return hash contains a `role` field which refers to the Organization Invitation role and will be one of the following values: `direct_member`, `admin`, `billing_manager`, `hiring_manager`, or `reinstate`. If the invitee is not a GitHub member, the `login` field in the return hash will be `null`.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/invitations`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-pending-invitations-legacy

https://developer.github.com/v3/teams/members/#list-pending-team-invitations-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List pending team invitations`](https://developer.github.com/v3/teams/members/#list-pending-team-invitations) endpoint.

The return hash contains a `role` field which refers to the Organization Invitation role and will be one of the following values: `direct_member`, `admin`, `billing_manager`, `hiring_manager`, or `reinstate`. If the invitee is not a GitHub member, the `login` field in the return hash will be `null`.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-projects-in-org

https://developer.github.com/v3/teams/#list-team-projects

Lists the organization projects for a team.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/projects`.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-projects-legacy

https://developer.github.com/v3/teams/#list-team-projects-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List team projects`](https://developer.github.com/v3/teams/#list-team-projects) endpoint.

Lists the organization projects for a team.

### parameters


| name | description |
|------|-------------|
| inertia-preview | __Required__ The Projects API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-10-27-changes-to-projects-api) for full details. To access the API during the preview period, you must set this flag. |
| team_id | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-repos-in-org

https://developer.github.com/v3/teams/#list-team-repositories

Lists a team's repositories visible to the authenticated user.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/repos`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams list-repos-legacy

https://developer.github.com/v3/teams/#list-team-repositories-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [List team repositories](https://developer.github.com/v3/teams/#list-team-repositories) endpoint.

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## teams remove-member-legacy

https://developer.github.com/v3/teams/members/#remove-team-member-legacy

The "Remove team member" endpoint (described below) is deprecated.

We recommend using the [Remove team membership for a user](https://developer.github.com/v3/teams/members/#remove-team-membership-for-a-user) endpoint instead. It allows you to remove both active and pending memberships.

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

To remove a team member, the authenticated user must have 'admin' permissions to the team or be an owner of the org that the team is associated with. Removing a team member does not delete the user, it just removes them from the team.

**Note:** When you have team synchronization set up for a team with your organization's identity provider (IdP), you will see an error if you attempt to use the API for making changes to the team's membership. If you have access to manage group membership in your IdP, you can manage GitHub team membership through your identity provider, which automatically adds and removes team members in an organization. For more information, see "[Synchronizing teams between your identity provider and GitHub](https://help.github.com/articles/synchronizing-teams-between-your-identity-provider-and-github/)."

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| username | __Required__  |

## teams remove-membership-for-user-in-org

https://developer.github.com/v3/teams/members/#remove-team-membership-for-a-user

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

To remove a membership between a user and a team, the authenticated user must have 'admin' permissions to the team or be an owner of the organization that the team is associated with. Removing team membership does not delete the user, it just removes their membership from the team.

**Note:** When you have team synchronization set up for a team with your organization's identity provider (IdP), you will see an error if you attempt to use the API for making changes to the team's membership. If you have access to manage group membership in your IdP, you can manage GitHub team membership through your identity provider, which automatically adds and removes team members in an organization. For more information, see "[Synchronizing teams between your identity provider and GitHub](https://help.github.com/articles/synchronizing-teams-between-your-identity-provider-and-github/)."

**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}/memberships/{username}`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| username | __Required__  |

## teams remove-membership-for-user-legacy

https://developer.github.com/v3/teams/members/#remove-team-membership-for-a-user-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Remove team membership for a user](https://developer.github.com/v3/teams/members/#remove-team-membership-for-a-user) endpoint.

Team synchronization is available for organizations using GitHub Enterprise Cloud. For more information, see [GitHub's products](https://help.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.

To remove a membership between a user and a team, the authenticated user must have 'admin' permissions to the team or be an owner of the organization that the team is associated with. Removing team membership does not delete the user, it just removes their membership from the team.

**Note:** When you have team synchronization set up for a team with your organization's identity provider (IdP), you will see an error if you attempt to use the API for making changes to the team's membership. If you have access to manage group membership in your IdP, you can manage GitHub team membership through your identity provider, which automatically adds and removes team members in an organization. For more information, see "[Synchronizing teams between your identity provider and GitHub](https://help.github.com/articles/synchronizing-teams-between-your-identity-provider-and-github/)."

### parameters


| name | description |
|------|-------------|
| team_id | __Required__  |
| username | __Required__  |

## teams remove-project-in-org

https://developer.github.com/v3/teams/#remove-a-project-from-a-team

Removes an organization project from a team. An organization owner or a team maintainer can remove any project from the team. To remove a project from a team as an organization member, the authenticated user must have `read` access to both the team and project, or `admin` access to the team or project. This endpoint removes the project from the team, but does not delete the project.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}/projects/{project_id}`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| project_id | __Required__  |
| team_slug | __Required__ team_slug parameter |

## teams remove-project-legacy

https://developer.github.com/v3/teams/#remove-a-project-from-a-team-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Remove a project from a team](https://developer.github.com/v3/teams/#remove-a-project-from-a-team) endpoint.

Removes an organization project from a team. An organization owner or a team maintainer can remove any project from the team. To remove a project from a team as an organization member, the authenticated user must have `read` access to both the team and project, or `admin` access to the team or project. **Note:** This endpoint removes the project from the team, but does not delete it.

### parameters


| name | description |
|------|-------------|
| project_id | __Required__  |
| team_id | __Required__  |

## teams remove-repo-in-org

https://developer.github.com/v3/teams/#remove-a-repository-from-a-team

If the authenticated user is an organization owner or a team maintainer, they can remove any repositories from the team. To remove a repository from a team as an organization member, the authenticated user must have admin access to the repository and must be able to see the team. This does not delete the repository, it just removes it from the team.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}/repos/{owner}/{repo}`.

### parameters


| name | description |
|------|-------------|
| org | __Required__  |
| repo | __Required__ repository in OWNER/REPO form |
| team_slug | __Required__ team_slug parameter |

## teams remove-repo-legacy

https://developer.github.com/v3/teams/#remove-a-repository-from-a-team-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Remove a repository from a team](https://developer.github.com/v3/teams/#remove-a-repository-from-a-team) endpoint.

If the authenticated user is an organization owner or a team maintainer, they can remove any repositories from the team. To remove a repository from a team as an organization member, the authenticated user must have admin access to the repository and must be able to see the team. NOTE: This does not delete the repository, it just removes it from the team.

### parameters


| name | description |
|------|-------------|
| repo | __Required__ repository in OWNER/REPO form |
| team_id | __Required__  |

## teams update-discussion-comment-in-org

https://developer.github.com/v3/teams/discussion_comments/#update-a-discussion-comment

Edits the body text of a discussion comment. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `PATCH /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.

### parameters


| name | description |
|------|-------------|
| body | __Required__ The discussion comment's body text. |
| comment_number | __Required__  |
| discussion_number | __Required__  |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams update-discussion-comment-legacy

https://developer.github.com/v3/teams/discussion_comments/#update-a-discussion-comment-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Update a discussion comment](https://developer.github.com/v3/teams/discussion_comments/#update-a-discussion-comment) endpoint.

Edits the body text of a discussion comment. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| body | __Required__ The discussion comment's body text. |
| comment_number | __Required__  |
| discussion_number | __Required__  |
| team_id | __Required__  |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |

## teams update-discussion-in-org

https://developer.github.com/v3/teams/discussions/#update-a-discussion

Edits the title and body text of a discussion post. Only the parameters you provide are updated. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

**Note:** You can also specify a team by `org_id` and `team_id` using the route `PATCH /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}`.

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| body | The discussion post's body text. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| title | The discussion post's title. |

## teams update-discussion-legacy

https://developer.github.com/v3/teams/discussions/#update-a-discussion-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Update a discussion](https://developer.github.com/v3/teams/discussions/#update-a-discussion) endpoint.

Edits the title and body text of a discussion post. Only the parameters you provide are updated. OAuth access tokens require the `write:discussion` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| discussion_number | __Required__  |
| team_id | __Required__  |
| body | The discussion post's body text. |
| squirrel-girl-preview | An additional `reactions` object in the issue comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2016-05-12-reactions-api-preview) for full details.<br><br>To access the API you must set this flag. |
| title | The discussion post's title. |

## teams update-in-org

https://developer.github.com/v3/teams/#update-a-team

To edit a team, the authenticated user must either be an organization owner or a team maintainer.

**Note:** You can also specify a team by `org_id` and `team_id` using the route `PATCH /organizations/{org_id}/team/{team_id}`.

### parameters


| name | description |
|------|-------------|
| name | __Required__ The name of the team. |
| org | __Required__  |
| team_slug | __Required__ team_slug parameter |
| description | The description of the team. |
| parent_team_id | The ID of a team to set as the parent team. |
| permission | **Deprecated**. The permission that new repositories will be added to the team with when none is specified. Can be one of:  <br>\* `pull` - team members can pull, but not push to or administer newly-added repositories.  <br>\* `push` - team members can pull and push, but not administer newly-added repositories.  <br>\* `admin` - team members can pull, push and administer newly-added repositories. |
| privacy | The level of privacy this team should have. Editing teams without specifying this parameter leaves `privacy` intact. When a team is nested, the `privacy` for parent teams cannot be `secret`. The options are:  <br>**For a non-nested team:**  <br>\* `secret` - only visible to organization owners and members of this team.  <br>\* `closed` - visible to all members of this organization.  <br>**For a parent or child team:**  <br>\* `closed` - visible to all members of this organization. |

## teams update-legacy

https://developer.github.com/v3/teams/#update-a-team-legacy

**Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Update a team](https://developer.github.com/v3/teams/#update-a-team) endpoint.

To edit a team, the authenticated user must either be an organization owner or a team maintainer.

**Note:** With nested teams, the `privacy` for parent teams cannot be `secret`.

### parameters


| name | description |
|------|-------------|
| name | __Required__ The name of the team. |
| team_id | __Required__  |
| description | The description of the team. |
| parent_team_id | The ID of a team to set as the parent team. |
| permission | **Deprecated**. The permission that new repositories will be added to the team with when none is specified. Can be one of:  <br>\* `pull` - team members can pull, but not push to or administer newly-added repositories.  <br>\* `push` - team members can pull and push, but not administer newly-added repositories.  <br>\* `admin` - team members can pull, push and administer newly-added repositories. |
| privacy | The level of privacy this team should have. Editing teams without specifying this parameter leaves `privacy` intact. The options are:  <br>**For a non-nested team:**  <br>\* `secret` - only visible to organization owners and members of this team.  <br>\* `closed` - visible to all members of this organization.  <br>**For a parent or child team:**  <br>\* `closed` - visible to all members of this organization. |

# users


## users add-email-for-authenticated

https://developer.github.com/v3/users/emails/#add-an-email-address-for-the-authenticated-user

This endpoint is accessible with the `user` scope.

## users block

https://developer.github.com/v3/users/blocking/#block-a-user



### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## users check-blocked

https://developer.github.com/v3/users/blocking/#check-if-a-user-is-blocked-by-the-authenticated-user

If the user is blocked:

If the user is not blocked:

### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## users check-following-for-user

https://developer.github.com/v3/users/followers/#check-if-a-user-follows-another-user



### parameters


| name | description |
|------|-------------|
| target_user | __Required__ target_user parameter |
| username | __Required__  |

## users check-person-is-followed-by-authenticated

https://developer.github.com/v3/users/followers/#check-if-a-person-is-followed-by-the-authenticated-user



### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## users create-gpg-key-for-authenticated

https://developer.github.com/v3/users/gpg_keys/#create-a-gpg-key-for-the-authenticated-user

Adds a GPG key to the authenticated user's GitHub account. Requires that you are authenticated via Basic Auth, or OAuth with at least `write:gpg_key` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| armored_public_key | __Required__ A GPG key in ASCII-armored format. |

## users create-public-ssh-key-for-authenticated

https://developer.github.com/v3/users/keys/#create-a-public-ssh-key-for-the-authenticated-user

Adds a public SSH key to the authenticated user's GitHub account. Requires that you are authenticated via Basic Auth, or OAuth with at least `write:public_key` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| key | __Required__ The public SSH key to add to your GitHub account. |
| title | A descriptive name for the new key. |

## users delete-email-for-authenticated

https://developer.github.com/v3/users/emails/#delete-an-email-address-for-the-authenticated-user

This endpoint is accessible with the `user` scope.

## users delete-gpg-key-for-authenticated

https://developer.github.com/v3/users/gpg_keys/#delete-a-gpg-key-for-the-authenticated-user

Removes a GPG key from the authenticated user's GitHub account. Requires that you are authenticated via Basic Auth or via OAuth with at least `admin:gpg_key` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| gpg_key_id | __Required__ gpg_key_id parameter |

## users delete-public-ssh-key-for-authenticated

https://developer.github.com/v3/users/keys/#delete-a-public-ssh-key-for-the-authenticated-user

Removes a public SSH key from the authenticated user's GitHub account. Requires that you are authenticated via Basic Auth or via OAuth with at least `admin:public_key` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| key_id | __Required__ key_id parameter |

## users follow

https://developer.github.com/v3/users/followers/#follow-a-user

Note that you'll need to set `Content-Length` to zero when calling out to this endpoint. For more information, see "[HTTP verbs](https://developer.github.com/v3/#http-verbs)."

Following a user requires the user to be logged in and authenticated with basic auth or OAuth with the `user:follow` scope.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## users get-authenticated

https://developer.github.com/v3/users/#get-the-authenticated-user

If the authenticated user is authenticated through basic authentication or OAuth with the `user` scope, then the response lists public and private profile information.

If the authenticated user is authenticated through OAuth without the `user` scope, then the response lists only public profile information.

## users get-by-username

https://developer.github.com/v3/users/#get-a-user

Provides publicly available information about someone with a GitHub account.

GitHub Apps with the `Plan` user permission can use this endpoint to retrieve information about a user's GitHub plan. The GitHub App must be authenticated as a user. See "[Identifying and authorizing users for GitHub Apps](https://developer.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/)" for details about authentication. For an example response, see "[Response with GitHub plan information](https://developer.github.com/v3/users/#response-with-github-plan-information)."

The `email` key in the following response is the publicly visible email address from your GitHub [profile page](https://github.com/settings/profile). When setting up your profile, you can select a primary email address to be “public” which provides an email entry for this endpoint. If you do not set a public email address for `email`, then it will have a value of `null`. You only see publicly visible email addresses when authenticated with GitHub. For more information, see [Authentication](https://developer.github.com/v3/#authentication).

The Emails API enables you to list all of your email addresses, and toggle a primary email to be visible publicly. For more information, see "[Emails API](https://developer.github.com/v3/users/emails/)".

### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## users get-context-for-user

https://developer.github.com/v3/users/#get-contextual-information-for-a-user

Provides hovercard information when authenticated through basic auth or OAuth with the `repo` scope. You can find out more about someone in relation to their pull requests, issues, repositories, and organizations.

The `subject_type` and `subject_id` parameters provide context for the person's hovercard, which returns more information than without the parameters. For example, if you wanted to find out more about `octocat` who owns the `Spoon-Knife` repository via cURL, it would look like this:

```shell
 curl -u username:token
  https://api.github.com/users/octocat/hovercard?subject_type=repository&subject_id=1300192
```

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| subject_id | Uses the ID for the `subject_type` you specified. **Required** when using `subject_type`. |
| subject_type | Identifies which additional information you'd like to receive about the person's hovercard. Can be `organization`, `repository`, `issue`, `pull_request`. **Required** when using `subject_id`. |

## users get-gpg-key-for-authenticated

https://developer.github.com/v3/users/gpg_keys/#get-a-gpg-key-for-the-authenticated-user

View extended details for a single GPG key. Requires that you are authenticated via Basic Auth or via OAuth with at least `read:gpg_key` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| gpg_key_id | __Required__ gpg_key_id parameter |

## users get-public-ssh-key-for-authenticated

https://developer.github.com/v3/users/keys/#get-a-public-ssh-key-for-the-authenticated-user

View extended details for a single public SSH key. Requires that you are authenticated via Basic Auth or via OAuth with at least `read:public_key` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| key_id | __Required__ key_id parameter |

## users list

https://developer.github.com/v3/users/#list-users

Lists all users, in the order that they signed up on GitHub. This list includes personal user accounts and organization accounts.

Note: Pagination is powered exclusively by the `since` parameter. Use the [Link header](https://developer.github.com/v3/#link-header) to get the URL for the next page of users.

### parameters


| name | description |
|------|-------------|
| per_page | Results per page (max 100) |
| since | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. |

## users list-blocked-by-authenticated

https://developer.github.com/v3/users/blocking/#list-users-blocked-by-the-authenticated-user

List the users you've blocked on your personal account.

## users list-emails-for-authenticated

https://developer.github.com/v3/users/emails/#list-email-addresses-for-the-authenticated-user

Lists all of your email addresses, and specifies which one is visible to the public. This endpoint is accessible with the `user:email` scope.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-followed-by-authenticated

https://developer.github.com/v3/users/followers/#list-the-people-the-authenticated-user-follows

Lists the people who the authenticated user follows.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-followers-for-authenticated-user

https://developer.github.com/v3/users/followers/#list-followers-of-the-authenticated-user

Lists the people following the authenticated user.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-followers-for-user

https://developer.github.com/v3/users/followers/#list-followers-of-a-user

Lists the people following the specified user.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-following-for-user

https://developer.github.com/v3/users/followers/#list-the-people-a-user-follows

Lists the people who the specified user follows.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-gpg-keys-for-authenticated

https://developer.github.com/v3/users/gpg_keys/#list-gpg-keys-for-the-authenticated-user

Lists the current user's GPG keys. Requires that you are authenticated via Basic Auth or via OAuth with at least `read:gpg_key` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-gpg-keys-for-user

https://developer.github.com/v3/users/gpg_keys/#list-gpg-keys-for-a-user

Lists the GPG keys for a user. This information is accessible by anyone.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-public-emails-for-authenticated

https://developer.github.com/v3/users/emails/#list-public-email-addresses-for-the-authenticated-user

Lists your publicly visible email address, which you can set with the [Set primary email visibility for the authenticated user](https://developer.github.com/v3/users/emails/#set-primary-email-visibility-for-the-authenticated-user) endpoint. This endpoint is accessible with the `user:email` scope.

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-public-keys-for-user

https://developer.github.com/v3/users/keys/#list-public-keys-for-a-user

Lists the _verified_ public SSH keys for a user. This is accessible by anyone.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users list-public-ssh-keys-for-authenticated

https://developer.github.com/v3/users/keys/#list-public-ssh-keys-for-the-authenticated-user

Lists the public SSH keys for the authenticated user's GitHub account. Requires that you are authenticated via Basic Auth or via OAuth with at least `read:public_key` [scope](https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).

### parameters


| name | description |
|------|-------------|
| page | Page number of the results to fetch. |
| per_page | Results per page (max 100) |

## users set-primary-email-visibility-for-authenticated

https://developer.github.com/v3/users/emails/#set-primary-email-visibility-for-the-authenticated-user

Sets the visibility for your primary email addresses.

### parameters


| name | description |
|------|-------------|
| email | __Required__ An email address associated with the GitHub user account to manage. |
| visibility | __Required__ Denotes whether an email is publically visible. |

## users unblock

https://developer.github.com/v3/users/blocking/#unblock-a-user



### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## users unfollow

https://developer.github.com/v3/users/followers/#unfollow-a-user

Unfollowing a user requires the user to be logged in and authenticated with basic auth or OAuth with the `user:follow` scope.

### parameters


| name | description |
|------|-------------|
| username | __Required__  |

## users update-authenticated

https://developer.github.com/v3/users/#update-the-authenticated-user

**Note:** If your email is set to private and you send an `email` parameter as part of this request to update your profile, your privacy settings are still enforced: the email address will not be displayed on your public profile or via the API.

### parameters


| name | description |
|------|-------------|
| bio | The new short biography of the user. |
| blog | The new blog URL of the user. |
| company | The new company of the user. |
| email | The publicly visible email address of the user. |
| hireable | The new hiring availability of the user. |
| location | The new location of the user. |
| name | The new name of the user. |
| twitter_username | The new Twitter username of the user. |
