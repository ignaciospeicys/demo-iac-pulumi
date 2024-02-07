### **1. Set Up AWS Account**

First, if you haven't already, create a personal AWS account. You'll then need to set up an IAM user or role with the necessary permissions for the resources you want to manage with Pulumi.

### **2. Download AWS CLI**

```bash
# For macOS with Homebrew:
brew install awscli

# For other platforms, follow the official AWS instructions:
# https://aws.amazon.com/cli/
```

### **3. Configure AWS Credentials**

After setting up your AWS account, configure your AWS credentials locally. You can do this in several ways:

- Create IAM User
- Create Access Key for that user (CLI use case)
- **AWS CLI**: Use the AWS CLI and run **`aws configure`** to set up your credentials. Pulumi will use these credentials by default.
### **4. Run Pulumi Up**

Run **`pulumi up`** to preview and deploy your AWS resources using your personal AWS account credentials.