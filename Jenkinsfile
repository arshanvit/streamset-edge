properties (
[
    [$class: 'GitLabConnectionProperty', gitLabConnection: 'gitlab'],
    [$class: 'ParametersDefinitionProperty', parameterDefinitions: [
        [$class: 'StringParameterDefinition', defaultValue: '', description: 'Some Description', name : 'RELEASE_VERSION']
        ]
    ]
])

pipeline = new org.qiotec.jenkins2.pipeline()

env.PROJECT_NAME = "datacollector-edge"
env.PROJECT_GROUP = "inference-analytics"
